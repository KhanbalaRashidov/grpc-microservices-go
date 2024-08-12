package ports

import (
	"context"
	"grpc-microservices-go/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
	GetOrder(ctx context.Context, id int64) (domain.Order, error)
}
