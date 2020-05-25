package main

import (
	"github.com/micro/go-micro"
	"gitlab.com/otis_team/backend/db/client"
	pb "gitlab.com/otis_team/backend/service/merchant/proto/merchant"
	"log"
)

/* Run the merchant service */
func main() {
	service := micro.NewService(
			micro.Name("go.micro.service.merchant"),
	)
	service.Init()

	dynamoClient := client.RDSClient{}
	var err error
	//err = dynamoClient.Init()
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
