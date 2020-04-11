package main

import (
		"github.com/micro/go-micro/v2"
		"gitlab.com/otis-team/backend/db/client"
		pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
		"log"
)

func main() {
	service := micro.NewService(
			micro.Name("go.micro.service.merchant"),
	)
	service.Init()

	dynamoClient := client.DynamoClient{}
	err := dynamoClient.Init()
	if err != nil {
		log.Panic(err)
	}

	handler := &Handler{dynamoClient}

	if err = pb.RegisterMerchantServiceHandler(service.Server(), handler); err != nil {
		log.Panic(err)
	}

	if err = service.Run(); err != nil {
		log.Println(err)
	}
}