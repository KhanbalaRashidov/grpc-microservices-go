package ports

import (
	"context"
	"grpc-microservices-go/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, domain domain.Order) (domain.Order, error)
	GetOrder(ctx context.Context, id int64) (domain.Order, error)
}
