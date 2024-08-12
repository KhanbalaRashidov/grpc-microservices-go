package ports

import (
	"context"
	"grpc-microservices-go/order/internal/application/core/domain"
)

type DbPort interface {
	Get(ctx context.Context, id int64) (domain.Order, error)
	Save(*domain.Order) error
}
