package main

import (
	"context"
	"fmt"
	"gitlab.com/otis-team/backend/service/transaction/package"
	"log"
	"os"
	"github.com/micro/go-micro"
	pb "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	
	service := micro.NewService(
			micro.Name("go.micro.service.transaction"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := transaction.CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	transactionCollection := client.Database("otis").Collection("transaction")

	repository := &transaction.MongoRepository{transactionCollection}
	h := &transaction.Handler{repository}

	pb.RegisterTransactionServiceHandler(service.Server(), h)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}