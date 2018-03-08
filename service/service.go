package service

import (
	pb "github.com/avegao/iot-openevse-service/resource/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"context"
	"github.com/avegao/openevse"
	"time"
	"github.com/avegao/gocondi"
	"github.com/avegao/openevse/command/ev_connect_state"
	"github.com/avegao/iot-openevse-service/entity/charger"
	"github.com/pkg/errors"
	"fmt"
	googleProtobuf "github.com/golang/protobuf/ptypes/empty"
)

type OpenevseService struct {
	pb.OpenevseServer
}

func (s OpenevseService) FindChargerById(ctx context.Context, request *pb.GetRequest) (*pb.Charger, error) {
	const logTag = "OpenevseService.FindChargerById"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	response := c.ToGrpcResponse()

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) FindAllChargers(ctx context.Context, request *googleProtobuf.Empty) (*pb.FindAllChargersResponse, error) {
	const logTag = "OpenevseService.FindAllChargers"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	chargers, err := charger.FindAll()
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if len(chargers) == 0 {
		err = errors.New("no chargers found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	var grpcChargers []*pb.Charger

	for _, c := range chargers {
		grpcChargers = append(grpcChargers, c.ToGrpcResponse())
	}

	response := &pb.FindAllChargersResponse{
		Chargers: grpcChargers,
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetAmmeterSettings(ctx context.Context, request *pb.GetRequest) (*pb.GetAmmeterSettingsResponse, error) {
	const logTag = "OpenevseService.GetAmmeterSettings"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	currentScaleFactor, currentOffset, err := openevse.GetAmmeterSettings(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetAmmeterSettingsResponse{
		CurrentScaleFactor: int32(currentScaleFactor),
		CurrentOffset:      int32(currentOffset),
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetAuthLockState(ctx context.Context, request *pb.GetRequest) (*pb.GetAuthLockStateResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetChargeLimit(ctx context.Context, request *pb.GetRequest) (*pb.GetChargeLimitResponse, error) {
	const logTag = "OpenevseService.GetChargeLimit"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	kwh, err := openevse.GetChargeLimit(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetChargeLimitResponse{
		Kwh: int32(kwh),
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetCurrentCapacityRangeInAmps(ctx context.Context, request *pb.GetRequest) (*pb.GetCurrentCapacityRangeInAmpsResponse, error) {
	const logTag = "OpenevseService.GetCurrentCapacityRangeInAmps"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	minAmps, maxAmps, err := openevse.GetCurrentCapacityRangeInAmps(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetCurrentCapacityRangeInAmpsResponse{
		MaxAmps: int32(maxAmps),
		MinAmps: int32(minAmps),
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetDelayTimer(ctx context.Context, request *pb.GetRequest) (*pb.GetDelayTimerResponse, error) {
	const logTag = "OpenevseService.GetDelayTimer"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	startTime, endTime, err := openevse.GetDelayTimer(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetDelayTimerResponse{
		StartTime: startTime,
		EndTime:   endTime,
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetEnergyUsage(ctx context.Context, request *pb.GetRequest) (*pb.GetEnergyUsageResponse, error) {
	const logTag = "OpenevseService.GetEnergyUsage"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	whInSession, kwhAccumulated, err := openevse.GetEnergyUsage(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetEnergyUsageResponse{
		WhInSession:    whInSession,
		KwhAccumulated: kwhAccumulated,
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetEvConnectState(ctx context.Context, request *pb.GetRequest) (*pb.GetEvConnectStateResponse, error) {
	const logTag = "OpenevseService.GetEvConnectState"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	state, err := openevse.GetEvConnectState(c.Host)

	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetEvConnectStateResponse{
		State: evStateToGrpc(state),
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetFaultCounters(ctx context.Context, request *pb.GetRequest) (*pb.GetFaultCountersResponse, error) {
	const logTag = "OpenevseService.GetFaultCounters"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	gfdi, noGround, stuckRelay, err := openevse.GetFaultCounters(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetFaultCountersResponse{
		Gfdi:       int32(gfdi),
		NoGround:   int32(noGround),
		StuckRelay: int32(stuckRelay),
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetOverTemperatureThresholds(ctx context.Context, request *pb.GetRequest) (*pb.GetOverTemperatureThresholdsResponse, error) {
	const logTag = "OpenevseService.GetOverTemperatureThresholds"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	ambient, ir, err := openevse.GetOverTemperatureThresholds(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetOverTemperatureThresholdsResponse{
		Ambient: ambient,
		Ir:      ir,
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetRtcTime(ctx context.Context, request *pb.GetRequest) (*pb.GetRtcTimeResponse, error) {
	const logTag = "OpenevseService.GetRtcTime"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	rtcTime, err := openevse.GetRtcTime(c.Host)

	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetRtcTimeResponse{
		RtcTime: rtcTime.Format(time.RFC3339),
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetSettings(ctx context.Context, request *pb.GetRequest) (*pb.GetSettingsResponse, error) {
	const logTag = "OpenevseService.GetSettings"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	amperes, flags, err := openevse.GetSettings(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetSettingsResponse{
		Amperes: int32(amperes),
		Flags:   []string{fmt.Sprintf("%v", flags)},
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetTimeLimit(ctx context.Context, request *pb.GetRequest) (*pb.GetTimeLimitResponse, error) {
	const logTag = "OpenevseService.GetTimeLimit"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	limit, err := openevse.GetTimeLimit(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetTimeLimitResponse{
		Limit: int32(limit),
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetVersion(ctx context.Context, request *pb.GetRequest) (*pb.GetVersionResponse, error) {
	const logTag = "OpenevseService.GetVersion"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	firmwareVersion, protocolVersion, err := openevse.GetVersion(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetVersionResponse{
		FirmwareVersion: firmwareVersion,
		ProtocolVersion: protocolVersion,
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) GetVoltmeterSettings(ctx context.Context, request *pb.GetRequest) (*pb.GetVoltmeterSettingsResponse, error) {
	const logTag = "OpenevseService.GetVoltmeterSettings"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	c, err := getChargerFromGetRequest(request)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	calefactor, offset, err := openevse.GetVoltmeterSettings(c.Host)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.GetVoltmeterSettingsResponse{
		Calefactor: int32(calefactor),
		Offset:     int32(offset),
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func (s OpenevseService) SetRtcTime(ctx context.Context, request *pb.SetRtcTimeRequest) (*pb.SetResponse, error) {
	const logTag = "OpenevseService.SetRtcTime"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("request", request).Debugf("%s - START", logTag)

	rtcTime, err := time.Parse(time.RFC3339, request.GetRtcTime())
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	c, err := charger.FindOneById(request.Id)
	if err != nil {
		logger.WithError(err).Errorf("%s - END", logTag)

		return nil, status.New(codes.Internal, err.Error()).Err()
	} else if c == nil {
		err = errors.New("charger not found")
		logger.WithError(err).Debugf("%s - END", logTag)

		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	if err := openevse.SetRtcTime(c.Host, rtcTime); err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &pb.SetResponse{
		Ok: true,
	}

	logger.WithField("response", response).Debugf("%s - END", logTag)

	return response, nil
}

func evStateToGrpc(state evConnectState.EvConnectState) (grpcState pb.GetEvConnectStateResponse_EvConnectState) {
	return pb.GetEvConnectStateResponse_EvConnectState(pb.GetEvConnectStateResponse_EvConnectState_value[state.String()])
}

func getChargerFromGetRequest(request *pb.GetRequest) (*charger.Charger, error) {
	return charger.FindOneById(request.GetId())
}
