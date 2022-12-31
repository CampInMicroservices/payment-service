package gapi

import (
	"context"
	"log"
	"payment-service/config"
	"payment-service/db"
	"payment-service/proto"
)

type ServerGRPC struct {
	config config.Config
	store  db.Store
	proto.UnimplementedPaymentServiceServer
}

func NewGrpcServer(config config.Config, store *db.Store) ServerGRPC {
	return ServerGRPC{
		config: config,
		store:  *store,
	}
}

func (s *ServerGRPC) CreatePaymentRequest(ctx context.Context, in *proto.PaymentRequest) (*proto.PaymentResponse, error) {
	log.Println("gRPC called! Creating new payment!")

	if in.Payment == nil {
		return nil, nil
	}

	arg := db.CreatePaymentParam{
		BookingID: in.Payment.BookingId,
		Price:     float64(in.Payment.Price),
		Paid:      in.Payment.Paid,
	}

	// Execute query.
	result, err := s.store.CreatePayment(ctx, arg)
	if err != nil {
		return nil, err
	}

	// Create response object
	response := &proto.PaymentResponse{
		Id: result.ID,
		Payment: &proto.Payment{
			BookingId: result.BookingID,
			Price:     float32(result.Price),
			Paid:      result.Paid,
		},
	}

	return response, nil
}
