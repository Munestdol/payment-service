package server

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	config "payment-service/configs"
	"payment-service/internal/service/proto"
	"payment-service/pkg/paymentservice"
)

type Server struct {
	server *http.Server
}

type ServerGRPC struct {
	server *grpc.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func NewServerGRPC() *ServerGRPC {
	return &ServerGRPC{grpc.NewServer()}
}

func (s *ServerGRPC) RegisterServices(services *proto.Service) {
	paymentservice.RegisterPaymentServiceServer(s.server, services)
}

func (s *ServerGRPC) Run(cfg *config.Config) error {
	fmt.Println(cfg.GRPC.Port)
	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatal().Err(err).Msg("error occurred while running grpc connection")

		return err
	}

	if err := s.server.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("error occurred while running grpc server")

		return err
	}

	return nil
}

func (s *ServerGRPC) Shutdown() {
	s.server.Stop()
}
