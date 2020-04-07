package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/errors"
	proto "gitlab.com/otis-team/backend/api/merchant/proto"
	"gitlab.com/otis-team/backend/db/model"
	merchantService "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	transactionService "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	"log"
)

// Merchant struct. All methods using this struct will be mapped to /merchant/<method>.
type Merchant struct {
	MerchantClient merchantService.MerchantServiceClient
}
// Transactions struct. All methods using this struct are mapped to /merchant/transaction/<method>
type Transactions struct {
	TransactionClient transactionService.TransactionService
}


// Create method (Merchant.Create) is served by HTTP requests to /merchant/create.
func (e *Merchant) Create(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
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

	var newMerchant *merchantService.Merchant
	err := json.Unmarshal([]byte(req.Body), &newMerchant)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", "Body not valid. Please reference to API documentation.")
	}

	r, err := e.MerchantClient.CreateMerchant(ctx, newMerchant)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"executed":   r.Created,
		"merchants": model.ProtobufToMerchant(r.Merchant),
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Get method (Merchant.Get) is served by HTTP requests to /merchant/get. Full endpoint is /merchant/get?id=<merchant_id>.
func (e *Merchant) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Merchant.Get request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.api.merchant", "Please provide an ID")
	}

	r, err := e.MerchantClient.GetMerchant(ctx, &merchantService.GetRequest{MerchantID: id.Values[0]})
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"executed": r.Merchants != nil,
		"merchant": model.ProtobufToMerchantCollection(r.Merchants),
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil

}

// GetAll method (Merchant.GetAll) is served by HTTP requests to /merchant/get-all.
func (e *Merchant) GetAll(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Merchant.GetAll request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	r, err := e.MerchantClient.GetMerchant(ctx, &merchantService.GetRequest{})
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"executed": r.Merchants != nil,
		"merchants": model.ProtobufToMerchantCollection(r.Merchants),
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Update method (Merchant.Update) is served by HTTP requests to /merchant/update.
func (e *Merchant) Update(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
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

	var updatedMerchant *model.Merchant
	err := json.Unmarshal([]byte(req.Body), &updatedMerchant)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", "Body not valid. Please reference to API documentation.")
	}

	r, err := e.MerchantClient.UpdateMerchant(ctx, model.MerchantToProtobuf(updatedMerchant))
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant",err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"executed":  r.Updated,
		"merchant": model.ProtobufToMerchant(r.Merchant),
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Delete : Method (Merchant.Delete) is served by HTTP requests to /merchant/delete. Full endpoint is /merchant/delete?id=<merchant_id>.
func (e *Merchant) Delete(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Merchant.Delete request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	merchantID, ok := req.Get["id"]
	if !ok || len(merchantID.Values) == 0 {
		return errors.BadRequest("go.micro.api.merchant", "Please provide an ID")
	}

	r, err := e.MerchantClient.DeleteMerchant(ctx, &merchantService.DeleteRequest{MerchantID: merchantID.Values[0]})
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"executed": r.Deleted,
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Get : Method (Transactions.Get) is served by HTTP requests to /merchant/transactions/get?id=<merchant_id> */
func (e *Transactions) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Merchant.Transactions request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	merchantID, ok := req.Get["id"]
	if !ok || len(merchantID.Values) == 0 {
		return errors.BadRequest("go.micro.api.merchant", "Please provide an ID")
	}

	r, err := e.TransactionClient.GetTransactions(ctx, &transactionService.IDRequest{MerchantID: merchantID.Values[0]})
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"executed": r.Transactions != nil,
		"transactions": model.ProtobufToTransactionCollection(r.Transactions),
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// Set JSON Body
	rsp.Body = string(body)

	return nil
}