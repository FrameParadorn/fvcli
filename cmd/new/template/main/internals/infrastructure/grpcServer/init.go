package grpcserver

import (
	grpcErrors "git.robodev.co/imp/shared-utility/grpc_errors"
	validatorUtils "git.robodev.co/imp/shared-utility/validator"
	"{{ service_name }}/internals/config"
	"{{ service_name }}/internals/controller"
	grpcHealthV1 "{{ service_name }}/pkg/grpc/health/v1"
	"google.golang.org/grpc"
)

type Server struct {
	Config           config.Configuration
	Server           *grpc.Server
	HealthCtrl       *controller.HealthZController
}

// Configure ...
func (s *Server) Configure() {
	grpcHealthV1.RegisterHealthServer(s.Server, s.HealthCtrl)
}

func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	validator *validatorUtils.CustomValidator,
) *Server {
	option := grpc.ChainUnaryInterceptor(
		grpcErrors.UnaryServerInterceptor(),
		validatorUtils.UnaryServerInterceptor(validator),
	)

	s := &Server{
		Server:           grpc.NewServer(option, grpc.MaxRecvMsgSize(10*10e6), grpc.MaxSendMsgSize(10*10e6)),
		Config:           config,
		HealthCtrl:       healthCtrl,
	}

	s.Configure()

	return s
}
