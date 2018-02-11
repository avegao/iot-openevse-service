package main

import (
	"github.com/avegao/openevse"
	"time"
	"os"
	"flag"
)

var host = flag.String("host", "", "Hostname")

func main() {
	flag.Parse()

	location, err := time.LoadLocation("Europe/Madrid")

	if err != nil {
		panic(err)
	}

	now := time.Now()
	now = now.In(location)

	if err := openevse.SetRtcTime(*host, now); err != nil {
		panic(err)
	}

	os.Exit(0)
}
