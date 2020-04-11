package main

import (
	"github.com/micro/go-micro/v2"
	"gitlab.com/otis-team/backend/db/client"
	pb "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	"log"
)

func main() {
	service := micro.NewService(
			micro.Name("go.micro.service.transaction"),
	)
	service.Init()

	dynamoClient := client.DynamoClient{}
	err := dynamoClient.Init()
	if err != nil {
		log.Panic(err)
	}

	handler := &Handler{dynamoClient}

	err = pb.RegisterTransactionServiceHandler(service.Server(), handler)
	if err != nil {
		log.Panic(err)
	}

	if err = service.Run(); err != nil {
		log.Println(err)
	}
}