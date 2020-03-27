package main

import (
	"context"
	"fmt"
	"gitlab.com/otis-team/backend/service/merchant/package"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	service := micro.NewService(
			micro.Name("go.micro.service.merchant"),
		)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := merchant.CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	merchantCollection := client.Database("otis").Collection("merchant")

	repository := &merchant.MongoRepository{merchantCollection}
	h := &merchant.Handler{repository}

	pb.RegisterMerchantServiceHandler(service.Server(), h)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}