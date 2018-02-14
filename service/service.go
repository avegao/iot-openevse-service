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

func (s OpenevseService) GetAmmeterSettings(ctx context.Context, request *pb.GetRequest) (*pb.GetAmmeterSettingsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetAuthLockState(ctx context.Context, request *pb.GetRequest) (*pb.GetAuthLockStateResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetChargeLimit(ctx context.Context, request *pb.GetRequest) (*pb.GetChargeLimitResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetCurrentCapacityRangeInAmps(ctx context.Context, request *pb.GetRequest) (*pb.GetCurrentCapacityRangeInAmpsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetDelayTimer(ctx context.Context, request *pb.GetRequest) (*pb.GetDelayTimerResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetEnergyUsage(ctx context.Context, request *pb.GetRequest) (*pb.GetEnergyUsageResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetEvConnectState(ctx context.Context, request *pb.GetRequest) (*pb.GetEvConnectStateResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetFaultCounters(ctx context.Context, request *pb.GetRequest) (*pb.GetFaultCountersResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetOverTemperatureThresholds(ctx context.Context, request *pb.GetRequest) (*pb.GetOverTemperatureThresholdsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetRtcTime(ctx context.Context, request *pb.GetRequest) (*pb.GetRtcTimeResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetSettings(ctx context.Context, request *pb.GetRequest) (*pb.GetSettingsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetTimeLimit(ctx context.Context, request *pb.GetRequest) (*pb.GetTimeLimitResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetVersion(ctx context.Context, request *pb.GetRequest) (*pb.GetVersionResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) GetVoltmeterSettings(ctx context.Context, request *pb.GetRequest) (*pb.GetVoltmeterSettingsResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

func (s OpenevseService) SetRtcTime(context.Context, *pb.SetRtcTimeRequest) (*pb.SetResponse, error) {
	return nil, status.New(codes.Unimplemented, "").Err()
}

