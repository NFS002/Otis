package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	protoAPI "gitlab.com/otis-team/backend/api/merchant/proto"
	merchant "gitlab.com/otis-team/backend/service/merchant/package"
	protoMerchant "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	"log"
)

type Merchant struct{
	client protoMerchant.MerchantServiceClient
}

type CreatedResponse struct {
	Created bool `json:"created"`
	MerchantID string `json:"merchantID"`
}

type GetResponse struct {
	Merchants []*merchant.Merchant `json:"merchants"`
}

type UpdateResponse struct{
	Updated bool `json:"update"`
	Merchant *merchant.Merchant `json:"merchant"`
}

type DeleteResponse struct{
	Deleted bool `json:"deleted"`
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

	var newMerchant *protoMerchant.Merchant
	err := json.Unmarshal([]byte(req.Body), &newMerchant)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", "Body not valid. Please reference to API documentation.")
	}

	r, err := e.client.CreateMerchant(ctx, newMerchant)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	rsp.StatusCode = 200

	createResponse := CreatedResponse{
		Created:   r.Created,
		MerchantID: r.MerchantID,
	}

	body, err := json.Marshal(createResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil

	return nil
}

func (e *Merchant) Get(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Get request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.api.merchant", "Please provide an ID")
	}

	r, err := e.client.GetMerchant(ctx, &protoMerchant.GetRequest{Id: id.Values[0]}) // Seems kinda janky
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	rsp.StatusCode = 200

	getResponse := GetResponse{
		Merchants: merchant.MarshalMerchantCollection(r.Merchants),
	}

	body, err := json.Marshal(getResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
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

	rsp.StatusCode = 200

	getResponse := GetResponse{
		Merchants: merchant.MarshalMerchantCollection(r.Merchants),
	}

	body, err := json.Marshal(getResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

func (e *Merchant) Update(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Update request")

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

	var updatedMerchant *merchant.Merchant
	err := json.Unmarshal([]byte(req.Body), &updatedMerchant)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", "Body not valid. Please reference to API documentation.")
	}

	r, err := e.client.UpdateMerchant(ctx, merchant.UnmarshalMerchant(updatedMerchant))
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	rsp.StatusCode = 200

	updateResponse := UpdateResponse{
		Updated:   r.Created,
		Merchant: merchant.MarshalMerchant(r.Merchant),
	}

	body, err := json.Marshal(updateResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

func (e *Merchant) Delete(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Delete request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	merchantID, ok := req.Get["id"]
	if !ok || len(merchantID.Values) == 0 {
		return errors.BadRequest("go.micro.api.merchant", "Please provide an ID")
	}

	r, err := e.client.DeleteMerchant(ctx, &protoMerchant.DeleteRequest{MerchantID: merchantID.Values[0]})
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	rsp.StatusCode = 200

	deleteResponse := DeleteResponse{Deleted: r.Deleted}

	body, err := json.Marshal(deleteResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
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
