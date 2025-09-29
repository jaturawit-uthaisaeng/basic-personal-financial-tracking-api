package grpc

import (
	protocol_buffer "basic-personal-financial-tracking-api/service/module/delivery/grpc/model"
	"basic-personal-financial-tracking-api/service/module/domain"

	"google.golang.org/grpc"
)

type newGrpcHandler struct {
	protocol_buffer.UnimplementedPersonalFinancialTrackingServiceServer
	usecase domain.PersonalFinancialTrackingUseCase
}

func NewServerGrpc(grpcServer *grpc.Server, usecase domain.PersonalFinancialTrackingUseCase) {
	protocol_buffer.RegisterPersonalFinancialTrackingServiceServer(grpcServer, &newGrpcHandler{
		usecase: usecase,
	})
}
