package ports

import "grpc-microservices-go/order/internal/application/core/domain"

type DbPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
