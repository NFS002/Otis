	package main

import (
	"log"
	"github.com/micro/go-micro"
	client "gitlab.com/otis-team/backend/db/client"
	pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	k8s "github.com/micro/examples/kubernetes/go/micro"

)

func main() {
	service := k8s.NewService(
			micro.Name("go.micro.service.merchant"),
	)
	service.Init()

	dynamoClient = client.DynamoClient{}
	err = dynamoClient.Init()
	if err != nil {
		log.Panic(err)
	}

	handler := &Handler{dynamoClient}

	pb.RegisterMerchantServiceHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}