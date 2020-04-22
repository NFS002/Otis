package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/errors"
	protoAPI "gitlab.com/otis_team/backend/api/merchant/proto"
	"gitlab.com/otis_team/backend/auth"
	merchant "gitlab.com/otis_team/backend/service/merchant/package"
	protoMerchant "gitlab.com/otis_team/backend/service/merchant/proto/merchant"

	transaction "gitlab.com/otis_team/backend/service/transaction/package"
	protoTransaction "gitlab.com/otis_team/backend/service/transaction/proto/transaction"
	"log"
)

// Merchant struct. All methods using this struct will be mapped to /merchant/<method>.
type Merchant struct {
	MerchantClient protoMerchant.MerchantServiceClient
}
// Transactions struct. All methods using this struct are mapped to /merchant/transaction/<method>
type Transactions struct {
	TransactionClient protoTransaction.TransactionService
}

// CreatedResponse maps CreateResponse protobuf message.
type CreatedResponse struct {
	Created bool `json:"created"`
	Merchant *merchant.Merchant `json:"merchantID"`
}

// GetResponse maps CreateResponse protobuf message.
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

// TransactionResponse maps TransactionResponse protobuf message.
type TransactionResponse struct {
	Transactions []*transaction.Transaction `json:"transactions"`
}

// Health method (Merchant.Health) is served by HTTP requests to /merchant/health.
func (e *Merchant) Health(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Health request")

	// Check request type
	if req.Method != "GET" {
		_ = errors.BadRequest(req.Url, "This method requires GET")
	}

	// Check auth
	authHeader, ok := req.Header["Authorization"]
	if !ok || len(authHeader.Values) == 0 {
		return errors.Unauthorized(req.Url, "Need Authorization header")
	}

	err := auth.CheckAuthHeader(authHeader.Values[0])
	if err != nil {
		return errors.Unauthorized(req.Url, err.Error())
	}

	// Return healthy
	type Health struct{
		Status string `json:"status"`
	}

	health := Health{Status: "OK"}

	body, err := json.Marshal(health)
	if err != nil {
		return errors.InternalServerError(req.Url, err.Error())
	}

	// set json body and code
	rsp.Body = string(body)
	rsp.StatusCode = 200

	return nil
}

// Create method (Merchant.Create) is served by HTTP requests to /merchant/create.
func (e *Merchant) Create(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Create request")

	// Check request type
	if req.Method != "POST" {
		return errors.BadRequest(req.Url, "This method requires POST")
	}

	// Check auth
	authHeader, ok := req.Header["Authorization"]
	if !ok || len(authHeader.Values) == 0 {
		return errors.Unauthorized(req.Url, "Need Authorization header")
	}

	err := auth.CheckAuthHeader(authHeader.Values[0])
	if err != nil {
		return errors.Unauthorized(req.Url, err.Error())
	}

	// Check content-type
	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest(req.Url, "Need Content-Type header")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest(req.Url, "Expect application/json")
	}

	// Create merchant
	var newMerchant *protoMerchant.Merchant
	err = json.Unmarshal([]byte(req.Body), &newMerchant)
	if err != nil {
		return errors.BadRequest(req.Url, "Body not valid. Please reference to API documentation.")
	}

	r, err := e.MerchantClient.CreateMerchant(ctx, newMerchant)
	if err != nil {
		return errors.InternalServerError(req.Url,err.Error())
	}

	// Create response
	createResponse := CreatedResponse{
		Created:   r.Created,
		Merchant: merchant.MarshalMerchant(r.Merchant),
	}

	body, err := json.Marshal(createResponse)
	if err != nil {
		return errors.InternalServerError(req.Url, err.Error())
	}

	// set json body and code
	rsp.Body = string(body)
	rsp.StatusCode = 200

	return nil
}

// Get method (Merchant.Get) is served by HTTP requests to /merchant/get. Full endpoint is /merchant/get?id=<merchant_id>.
func (e *Merchant) Get(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Get request")

	// Check request type
	if req.Method != "GET" {
		return errors.BadRequest(req.Url, "This method requires GET")
	}

	// Check auth
	authHeader, ok := req.Header["Authorization"]
	if !ok || len(authHeader.Values) == 0 {
		return errors.Unauthorized(req.Url, "Need Authorization header")
	}

	err := auth.CheckAuthHeader(authHeader.Values[0])
	if err != nil {
		return errors.Unauthorized(req.Url, err.Error())
	}

	// Check id
	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest(req.Url, "Please provide an ID")
	}

	// Get merchant
	r, err := e.MerchantClient.GetMerchant(ctx, &protoMerchant.GetRequest{MerchantID: id.Values[0]}) // Seems kinda janky
	if err != nil {
		return errors.InternalServerError(req.Url,err.Error())
	}

	// Create response
	getResponse := GetResponse{
		Merchants: merchant.MarshalMerchantCollection(r.Merchants),
	}

	body, err := json.Marshal(getResponse)
	if err != nil {
		return errors.InternalServerError(req.Url, err.Error())
	}

	// set json body
	rsp.Body = string(body)
	rsp.StatusCode = 200

	return nil
}

// GetAll method (Merchant.GetAll) is served by HTTP requests to /merchant/get-all.
func (e *Merchant) GetAll(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.GetAll request")

	// Check request type
	if req.Method != "GET" {
		return errors.BadRequest(req.Url, "This method requires GET")
	}

	// Check auth
	authHeader, ok := req.Header["Authorization"]
	if !ok || len(authHeader.Values) == 0 {
		return errors.Unauthorized(req.Url, "Need Authorization header")
	}

	err := auth.CheckAuthHeader(authHeader.Values[0])
	if err != nil {
		return errors.Unauthorized(req.Url, err.Error())
	}

	// Get all merchants
	r, err := e.MerchantClient.GetMerchant(ctx, &protoMerchant.GetRequest{})
	if err != nil {
		return errors.InternalServerError(req.Url,err.Error())
	}

	// Create response
	getResponse := GetResponse{
		Merchants: merchant.MarshalMerchantCollection(r.Merchants),
	}

	body, err := json.Marshal(getResponse)
	if err != nil {
		return errors.InternalServerError(req.Url, err.Error())
	}

	// set json body and code
	rsp.Body = string(body)
	rsp.StatusCode = 200

	return nil
}

// Update method (Merchant.Update) is served by HTTP requests to /merchant/update.
func (e *Merchant) Update(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Update request")

	// Check request type
	if req.Method != "POST" {
		return errors.BadRequest(req.Url, "This method requires POST")
	}

	// Check auth
	authHeader, ok := req.Header["Authorization"]
	if !ok || len(authHeader.Values) == 0 {
		return errors.Unauthorized(req.Url, "Need Authorization header")
	}

	err := auth.CheckAuthHeader(authHeader.Values[0])
	if err != nil {
		return errors.Unauthorized(req.Url, err.Error())
	}
	
	// Check content type
	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest(req.Url, "Need Content-Type header")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest(req.Url, "Expect application/json")
	}

	// Update merchant
	var updatedMerchant *merchant.Merchant
	err = json.Unmarshal([]byte(req.Body), &updatedMerchant)
	if err != nil {
		return errors.BadRequest(req.Url, "Body not valid. Please reference to API documentation.")
	}

	r, err := e.MerchantClient.UpdateMerchant(ctx, merchant.UnmarshalMerchant(updatedMerchant))
	if err != nil {
		return errors.InternalServerError(req.Url,err.Error())
	}
	
	// Create response
	updateResponse := UpdateResponse{
		Updated:  r.Updated,
		Merchant: merchant.MarshalMerchant(r.Merchant),
	}

	body, err := json.Marshal(updateResponse)
	if err != nil {
		return errors.InternalServerError(req.Url, err.Error())
	}

	// set json body
	rsp.Body = string(body)
	rsp.StatusCode = 200

	return nil
}

// Delete : Method (Merchant.Delete) is served by HTTP requests to /merchant/delete. Full endpoint is /merchant/delete?id=<merchant_id>.
func (e *Merchant) Delete(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Delete request")

	// Check request type
	if req.Method != "GET" {
		return errors.BadRequest(req.Url, "This method requires GET")
	}

	// Check auth
	authHeader, ok := req.Header["Authorization"]
	if !ok || len(authHeader.Values) == 0 {
		return errors.Unauthorized(req.Url, "Need Authorization header")
	}

	err := auth.CheckAuthHeader(authHeader.Values[0])
	if err != nil {
		return errors.Unauthorized(req.Url, err.Error())
	}

	// Check id
	merchantID, ok := req.Get["id"]
	if !ok || len(merchantID.Values) == 0 {
		return errors.BadRequest(req.Url, "Please provide an ID")
	}

	// Delete merchant
	r, err := e.MerchantClient.DeleteMerchant(ctx, &protoMerchant.DeleteRequest{MerchantID: merchantID.Values[0]})
	if err != nil {
		return errors.InternalServerError(req.Url, err.Error())
	}

	// Create response
	deleteResponse := DeleteResponse{Deleted: r.Deleted}
	body, err := json.Marshal(deleteResponse)
	if err != nil {
		return errors.InternalServerError(req.Url, err.Error())
	}

	// set json body and code
	rsp.Body = string(body)
	rsp.StatusCode = 200

	return nil
}

// Get : Method (Transactions.Get) is served by HTTP requests to /merchant/transactions/get?id=<merchant_id> */
func (e *Transactions) Get(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received Merchant.Transactions request")

	// Check request type
	if req.Method != "GET" {
		return errors.BadRequest(req.Url, "This method requires GET")
	}

	// Check auth
	authHeader, ok := req.Header["Authorization"]
	if !ok || len(authHeader.Values) == 0 {
		return errors.Unauthorized(req.Url, "Need Authorization header")
	}

	err := auth.CheckAuthHeader(authHeader.Values[0])
	if err != nil {
		return errors.Unauthorized(req.Url, err.Error())
	}

	// Check id
	merchantID, ok := req.Get["id"]
	if !ok || len(merchantID.Values) == 0 {
		return errors.BadRequest(req.Url, "Please provide an ID")
	}

	// Get transactions
	r, err := e.TransactionClient.GetTransactions(ctx, &protoTransaction.IDRequest{MerchantID: merchantID.Values[0]})
	if err != nil {
		return errors.InternalServerError(req.Url, err.Error())
	}

	// Create response
	transactionResponse := TransactionResponse{Transactions: transaction.MarshalTransactionCollection(r.Transactions)}
	body, err := json.Marshal(transactionResponse)
	if err != nil {
		return errors.InternalServerError(req.Url, err.Error())
	}

	// Set JSON Body and code
	rsp.Body = string(body)
	rsp.StatusCode = 200

	return nil
}
