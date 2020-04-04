package main

import (
	"log"
	client "gitlab.com/otis-team/backend/db/client"
	"github.com/micro/go-micro"
	pb "gitlab.com/otis-team/backend/service/user/proto/user"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
	)
	service.Init()

	dynamoClient = client.DynamoClient{}
	dynamoClient.Init()

	handler := &Handler{dynamoClient}

	pb.RegisterUserServiceHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}