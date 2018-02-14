package service

import (
	pb "github.com/avegao/iot-openevse-service/resource/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"context"
)

type OpenevseService struct {
	pb.OpenevseServer
}

func (s OpenevseService) GetAmmeterSettings(context.Context, *pb.GetRequest) (*pb.GetAmmeterSettingsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetAuthLockState(context.Context, *pb.GetRequest) (*pb.GetAuthLockStateResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetChargeLimit(context.Context, *pb.GetRequest) (*pb.GetChargeLimitResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetCurrentCapacityRangeInAmps(context.Context, *pb.GetRequest) (*pb.GetCurrentCapacityRangeInAmpsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetDelayTimer(context.Context, *pb.GetRequest) (*pb.GetDelayTimerResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetEnergyUsage(context.Context, *pb.GetRequest) (*pb.GetEnergyUsageResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetEvConnectState(context.Context, *pb.GetRequest) (*pb.GetEvConnectStateResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetFaultCounters(context.Context, *pb.GetRequest) (*pb.GetFaultCountersResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetOverTemperatureThresholds(context.Context, *pb.GetRequest) (*pb.GetOverTemperatureThresholdsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetRtcTime(context.Context, *pb.GetRequest) (*pb.GetRtcTimeResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetSettings(context.Context, *pb.GetRequest) (*pb.GetSettingsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetTimeLimit(context.Context, *pb.GetRequest) (*pb.GetTimeLimitResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetVersion(context.Context, *pb.GetRequest) (*pb.GetVersionResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetVoltmeterSettings(context.Context, *pb.GetRequest) (*pb.GetVoltmeterSettingsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) SetRtcTime(context.Context, *pb.SetRtcTimeRequest) (*pb.SetResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

