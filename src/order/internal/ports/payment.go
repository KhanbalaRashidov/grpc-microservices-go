package ports

import "grpc-microservices-go/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}
