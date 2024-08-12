package ports

import (
	"context"
	"grpc-microservices-go/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
