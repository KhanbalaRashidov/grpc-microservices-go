package api

import (
	"grpc-microservices-go/order/internal/application/core/domain"
	"grpc-microservices-go/order/internal/ports"
)

type Application struct {
	db      ports.DbPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DbPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (app Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := app.db.Save(order)
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := app.payment.Charge(&order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}

	return order, nil
}
