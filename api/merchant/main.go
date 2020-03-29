package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	protoAPI "gitlab.com/otis-team/backend/api/merchant/proto"

	protoMerchant "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	protoTransaction "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	"log"
)


func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.merchant"),
	)

	service.Init()

	mClient := protoMerchant.NewMerchantServiceClient("go.micro.service.merchant", service.Client())
	tClient := protoTransaction.NewTransactionServiceClient("go.micro.service.transaction", service.Client())

	handler := &Merchant{MerchantClient: mClient, TransactionClient: tClient}

	// Registering merchant API handler
	err := protoAPI.RegisterMerchantHandler(service.Server(), handler)
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
