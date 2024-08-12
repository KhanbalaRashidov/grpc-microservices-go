package grpc

import (
	"context"
	"fmt"
	"github.com/KhanbalaRashidov/microservices-proto/golang/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-microservices-go/payment/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}
