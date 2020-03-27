package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	merchant "gitlab.com/otis-team/backend/service/merchant/package"
	protoMerchant "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	protoAPI "gitlab.com/otis-team/backend/api/merchant/proto"
	"log"
)

type Merchant struct{
	client protoMerchant.MerchantServiceClient
}

// Merchant.Create is a method which will be served by http request /merchant/create
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /merchant/Create goes to go.micro.api.merchant Merchant.Create

func (e *Merchant) Create(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Create request")

	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires POST")
	}

	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.merchant", "Need Content-Type header")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.example", "Expect application/json")
	}

	var merchant *protoMerchant.Merchant
	err := json.Unmarshal([]byte(req.Body), &merchant)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", "Body not valid. Please reference to API documentation.")
	}

	r, err := e.client.CreateMerchant(ctx, merchant)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	// set response status
	rsp.StatusCode = 200

	// respond with some json - Surely can unmarshal to JSON
	var created string
	if r.Created == true {
		created = "true"
	} else {
		created = "false"
	}

	b, _ := json.Marshal(map[string]string{
		"created": created,
		"merchant_id": r.MerchantID,
	})

	log.Print("Received Merchant.Create request")

	// set json body
	rsp.Body = string(b)

	return nil
}

func (e *Merchant) Get(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Get request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.merchant", "Please provide an ID")
	}

	r, err := e.client.GetMerchant(ctx, &protoMerchant.GetRequest{Id: id.Values[0]}) // Seems kinda janky
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	merchantCollection := merchant.MarshalMerchantCollection(r.Merchants)
	body, err := json.Marshal(merchantCollection)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	rsp.StatusCode = 200
	rsp.Body = string(body)

	return nil
}

func (e *Merchant) GetAll(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.GetAll request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	r, err := e.client.GetMerchant(ctx, &protoMerchant.GetRequest{})
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	merchantCollection := merchant.MarshalMerchantCollection(r.Merchants)
	body, err := json.Marshal(merchantCollection)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	rsp.StatusCode = 200
	rsp.Body = string(body)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.merchant"),
	)

	service.Init()

	client := protoMerchant.NewMerchantServiceClient("go.micro.service.merchant", service.Client())

	m := &Merchant{client}

	// Registering merchant api handler
	err := protoAPI.RegisterMerchantHandler(service.Server(), m)
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
