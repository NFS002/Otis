package main

import (
	"github.com/micro/go-micro/v2"
	proto "gitlab.com/otis_team/backend/api/merchant/proto"
	merchantService "gitlab.com/otis_team/backend/service/merchant/proto/merchant"
	transactionService "gitlab.com/otis_team/backend/service/transaction/proto/transaction"
	"log"
)


/* This service sits inbetween the API gateway and the user/merchant/transaction services.
 * The main function of this package receives rpc requests forwarded from the API gateway, 
 * and then forward these to either the merchant and transaction servic, returning the response
 * to the API */
func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.merchant"),
	)

	service.Init()


	merchantClient := merchantService.NewMerchantService("go.micro.service.merchant", service.Client())
	
	transactionClient := transactionService.NewTransactionService("go.micro.service.transaction", service.Client())

	merchantHandler := &Merchant{ MerchantClient: merchantClient }
	transactionHandler := &Transactions{ TransactionClient: transactionClient }

	// Registering both API handlers

	merchantErr := proto.RegisterMerchantHandler( service.Server(), merchantHandler )
	if merchantErr != nil {
		log.Fatal(merchantErr)
	}

	transactionErr := proto.RegisterTransactionHandler(service.Server(), transactionHandler )
	if transactionErr != nil {
		log.Fatal(transactionErr)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
