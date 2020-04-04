package main

import (
	"log"
	"github.com/micro/go-micro"
	client "gitlab.com/otis-team/backend/db/client"
	pb "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
)

func main() {
	service := micro.NewService(
			micro.Name("go.micro.service.transaction"),
	)
	service.Init()

	dynamoClient = client.DynamoClient{}
	err = dynamoClient.Init()
	if err != nil {
		log.Panic(err)
	}

	handler := &Handler{dynamoClient}

	pb.RegisterTransactionServiceHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}