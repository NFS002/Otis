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

// Merchant struct. All methods using this struct will be mapped to /merchant/<method>.
type Merchant struct{
	client protoMerchant.MerchantServiceClient
}

// CreatedResponse maps CreateResponse protobuf message.
type CreatedResponse struct {
	Created bool `json:"created"`
	Merchant *merchant.Merchant `json:"merchantID"`
}

// CreatedResponse maps CreateResponse protobuf message.
type GetResponse struct {
	Merchants []*merchant.Merchant `json:"merchants"`
}

// UpdateResponse maps UpdateResponse protobuf message.
type UpdateResponse struct{
	Updated bool `json:"update"`
	Merchant *merchant.Merchant `json:"merchant"`
}

// DeleteResponse maps DeleteResponse protobuf message.
type DeleteResponse struct{
	Deleted bool `json:"deleted"`
}


// Create method (Merchant.Create) is served by HTTP requests to /merchant/create.
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
		Merchant: merchant.MarshalMerchant(r.Merchant),
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

// Get method (Merchant.Get) is served by HTTP requests to /merchant/get. Full endpoint is /merchant/get?id=<merchant_id>.
func (e *Merchant) Get(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Get request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.api.merchant", "Please provide an ID")
	}

	r, err := e.client.GetMerchant(ctx, &protoMerchant.GetRequest{MerchantID: id.Values[0]}) // Seems kinda janky
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

// GetAll method (Merchant.GetAll) is served by HTTP requests to /merchant/get-all.
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

// Update method (Merchant.Update) is served by HTTP requests to /merchant/update.
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
		Updated:   r.Updated,
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

// Delete method (Merchant.Delete) is served by HTTP requests to /merchant/delete. Full endpoint is /merchant/delete?id=<merchant_id>.
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
