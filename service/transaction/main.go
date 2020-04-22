package main

import (
	"context"
	"fmt"
	"gitlab.com/otis_team/backend/service/transaction/package"
	"log"
	"os"
	"github.com/micro/go-micro"
	pb "gitlab.com/otis_team/backend/service/transaction/proto/transaction"
)

const (
	defaultHost = "datastore:27017"
)

/* Connect to a MongoDB instance, and start and run a new grpc server on a Microservice,
 * passing that connection to the handler functions.  */
func main() {
	f := transaction.Transaction{ MerchantID: "44 "}
	log.Printf( "Transaction: %v",f)
	
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

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
