package main

import (
	"flag"
	"github.com/avegao/gocondi"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/heroku/rollrus"
	"os"
	"os/signal"
	"syscall"
	"google.golang.org/grpc"
	"net"
	"google.golang.org/grpc/reflection"
	pb "github.com/avegao/iot-openevse-service/resource/grpc"
	"github.com/avegao/iot-openevse-service/service"
)

const version = "1.0.0"

var (
	debug      = flag.Bool("debug", false, "Print debug logs")
	grcpPort   = flag.Int("port", 50000, "gRPC Server port. Default = 50000")
	buildDate  string
	commitHash string
	container  *gocondi.Container
	server     *grpc.Server
	parameters map[string]interface{}
)

func initContainer() {
	flag.Parse()

	parameters = map[string]interface{}{
		"build_date":              buildDate,
		"debug":                   *debug,
		"commit_hash":             commitHash,
		"version":                 version,
	}

	logger := initLogger()
	gocondi.Initialize(logger)
	container = gocondi.GetContainer()

	for name, value := range parameters {
		container.SetParameter(name, value)
	}
}

func initLogger() *logrus.Logger {
	logLevel := logrus.InfoLevel
	environment := "release"
	log := logrus.New()

	if *debug {
		logLevel = logrus.DebugLevel
		environment = "debug"
	} else {
		hook := rollrus.NewHook(fmt.Sprintf("%v", parameters["rollbar_token"]), environment)
		log.Hooks.Add(hook)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
	log.SetLevel(logLevel)

	return log
}

func initGrpc() {
	container.GetLogger().Debugf("initGrpc() - START")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *grcpPort))

	if err != nil {
		container.GetLogger().Fatalf("failed to listen: %v", err)
	}

	container.GetLogger().Debugf("gRPC listening in %d port", *grcpPort)

	server = grpc.NewServer()
	pb.RegisterOpenevseServer(server, new(service.OpenevseService))
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		container.GetLogger().Fatalf("failed to server: %v", err)
	}

	container.GetLogger().Debugf("initGrpc() - END")
}

func handleInterrupt() {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		<-gracefulStop
		powerOff()
	}()
}

func powerOff() {
	container.GetLogger().Infof("Shutting down...")

	if server != nil {
		server.Stop()
	}

	if container != nil {
		container.Close()
	}

	os.Exit(0)
}

func main() {
	initContainer()
	handleInterrupt()
	defer powerOff()

	logger := container.GetLogger()
	logger.Infof("IoT OpenEVSE Service started v%s (commit %s, build date %s)", container.GetStringParameter("version"), container.GetStringParameter("commit_hash"), container.GetStringParameter("build_date"))

	initGrpc()
}
