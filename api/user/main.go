package main

import (
	"github.com/micro/go-micro"
	protoAPI "gitlab.com/otis-team/backend/api/user/proto"

	protoUser "gitlab.com/otis-team/backend/service/user/proto/user"
	protoTransaction "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	"log" 
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	service.Init()

	userClient := protoUser.NewUserServiceClient("go.micro.service.user", service.Client())
	transactionClient := protoTransaction.NewTransactionService("go.micro.service.transaction", service.Client())


	userHandler := &User{ Client: userClient }
	transactionHandler := &Transactions{ Client: transactionClient }


	// Registering both API handlers
	userErr := protoAPI.RegisterUserHandler(service.Server(), userHandler )
	if userErr != nil {
		log.Fatal(userErr)
	}

	transactionErr := protoAPI.RegisterTransactionHandler( service.Server(), transactionHandler )
	if transactionErr != nil {
		log.Fatal(transactionErr)
	}


	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
