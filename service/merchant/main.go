	package main

import (
	"log"
	"github.com/micro/go-micro"
	client "gitlab.com/otis-team/backend/db/client"
	pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"

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
	pb.RegisterMerchantServiceHandler(service.Server(), handler)
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}