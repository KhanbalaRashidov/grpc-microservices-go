package main

import (
	"grpc-microservices-go/order/config"
	"grpc-microservices-go/order/internal/adapters/db"
	"grpc-microservices-go/order/internal/adapters/grpc"
	"grpc-microservices-go/order/internal/adapters/payment"
	"grpc-microservices-go/order/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSource())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
