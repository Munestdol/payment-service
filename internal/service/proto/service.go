package proto

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"payment-service/pkg/paymentservice"
)

type Payment interface {
	ChangeStatus(ctx context.Context, order *paymentservice.PaymentResult) (*emptypb.Empty, error)
}

type Service struct {
	Payment
	paymentservice.UnsafePaymentServiceServer
}

//func NewService(repo *proto.Repository) *Service {
//	return &Service{
//		Order:      NewOrderService(repo),
//		Manager:    NewManagerService(repo),
//		Basket:     NewBasketService(repo),
//		Restaurant: NewRestaurantService(repo),
//	}
//}
