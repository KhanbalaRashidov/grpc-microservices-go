package main

import (
	"grpc-microservices-go/payment/config"
	"grpc-microservices-go/payment/internal/adapters/db"
	"grpc-microservices-go/payment/internal/adapters/grpc"
	"grpc-microservices-go/payment/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSource())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
