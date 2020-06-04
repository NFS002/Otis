package main

import (
	"github.com/micro/go-micro"
	"gitlab.com/otis_team/backend/db/client"
	pb "gitlab.com/otis_team/backend/service/user/proto/user"
	"log"
)

/* Run the user service */
func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
	)
	service.Init()

	dynamoClient := client.RDSClient{}
	var err error
	//err = dynamoClient.Init()
	if err != nil {
		log.Panic(err)
	}

	handler := &Handler{dynamoClient}

	err = pb.RegisterUserServiceHandler(service.Server(), handler)
	if err != nil {
		log.Panic(err)
	}

	if err = service.Run(); err != nil {
		log.Println(err)
	}
}
