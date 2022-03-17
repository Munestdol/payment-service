package service

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"payment-service/pkg/paymentservice"
	"time"
)

func (s *PaymentService) CreateConnectionFD() (paymentservice.PaymentServiceClient, *grpc.ClientConn, context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", s.cfg.GRPCFD.Host, s.cfg.GRPCFD.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error().Err(err).Msg("error occurred while creating conn to FD")
		fmt.Println(err)
		return nil, nil, ctx, err
	}

	orderUpdateClient := paymentservice.NewPaymentServiceClient(conn)

	return orderUpdateClient, conn, ctx, nil
}

//make con to RA and test all shit
//func (s *PaymentService) CreateConnectionRA() (paymentservice.PaymentServiceClient, *grpc.ClientConn, context.Context, error) {
//	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
//	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", s.cfg.GRPCRA.Host, s.cfg.GRPCRA.Port), grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		log.Error().Err(err).Msg("error occurred while creating conn to FD")
//
//		return nil, nil, ctx, err
//	}
//
//	orderUpdateClient := paymentservice.NewPaymentServiceClient(conn)
//
//	return orderUpdateClient, conn, ctx, nil
//}

func (s *PaymentService) ChangeStatusFD(answer bool, id string) error {

	orderClientFD, conn, ctx, err := s.CreateConnectionFD()
	if err != nil {
		return err
	}

	updateOrder := &paymentservice.PaymentResult{
		Answer:  answer,
		IdOrder: id,
	}

	if _, err = orderClientFD.ChangeStatus(ctx, updateOrder); err != nil {
		log.Error().Err(err).Msg("error occurred while updating order in FD")
		return err
	}
	defer conn.Close()
	return nil
}
