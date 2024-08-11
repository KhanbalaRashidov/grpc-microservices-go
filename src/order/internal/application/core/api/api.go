package api

import (
	"grpc-microservices-go/order/internal/application/core/domain"
	"grpc-microservices-go/order/internal/ports"
)

type Application struct {
	db ports.DbPort
}

func NewApplication(db ports.DbPort) *Application {
	return &Application{
		db: db,
	}
}

func (app Application) PlaceOrder(oder domain.Order) (domain.Order, error) {
	err := app.db.Save(oder)
	if err != nil {
		return domain.Order{}, err
	}
	return oder, nil
}
