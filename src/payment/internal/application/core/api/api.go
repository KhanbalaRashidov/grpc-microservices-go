package api

import (
	"context"
	"grpc-microservices-go/payment/internal/application/core/domain"
	"grpc-microservices-go/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(dbPort ports.DBPort) *Application {
	return &Application{
		db: dbPort,
	}
}

func (app Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	err := app.db.Save(ctx, &payment)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}
