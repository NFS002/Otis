package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/errors"
	protoAPI "gitlab.com/otis-team/backend/api/user/proto"

	protoUser "gitlab.com/otis-team/backend/service/user/proto/user"
	user "gitlab.com/otis-team/backend/service/user/package"

	protoTransaction "gitlab.com/otis-team/backend/service/transaction/proto/transaction"
	transaction "gitlab.com/otis-team/backend/service/transaction/package"
	"log"
)

// User struct. All methods using this struct will be mapped to /user/<method>.
type User struct {
	Client protoUser.UserServiceClient
}

// Transactions struct. All methods using this struct will be mapped to /user/transaction/<method>
type Transactions struct {
	Client protoTransaction.TransactionService
}

// CreatedResponse maps CreateResponse protobuf message.
type CreatedResponse struct {
	Created bool `json:"created"`
	User *user.User `json:"user"`
}

// GetResponse maps CreateResponse protobuf message.
type GetResponse struct {
	Users []*user.User `json:"users"`
}

// UpdateResponse maps UpdateResponse protobuf message.
type UpdateResponse struct{
	Updated bool `json:"update"`
	User *user.User `json:"user"`
}

// DeleteResponse maps DeleteResponse protobuf message.
type DeleteResponse struct{
	Deleted bool `json:"deleted"`
}

// TransactionResponse maps TransactionResponse protobuf message.
type TransactionResponse struct {
	Transactions []*transaction.Transaction `json:"transactions"`
} 


// Create method (User.Create) is served by HTTP requests to /user/create.
func (e *User) Create(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received User.Create request")

	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.user", "This method requires POST")
	}

	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Need Content-Type header")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.example", "Expect application/json")
	}

	var newUser *user.User
	err := json.Unmarshal([]byte(req.Body), &newUser)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", "Body not valid. Please reference to API documentation.")
	}

	r, err := e.Client.CreateUser(ctx, user.UnmarshalUser(newUser))
	if err != nil {
		return errors.BadRequest("go.micro.api.user",err.Error())
	}

	rsp.StatusCode = 200

	createResponse := CreatedResponse{
		Created:   r.Created,
		User: user.MarshalUser(r.User),
	}

	body, err := json.Marshal(createResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Get method (User.Get) is served by HTTP requests to /user/get. Full endpoint is /user/get?id=<user_id>.
func (e *User) Get(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received User.Get request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Please provide an ID")
	}

	r, err := e.Client.GetUser(ctx, &protoUser.GetRequest{UserID: id.Values[0]}) // Seems kinda janky
	if err != nil {
		return errors.BadRequest("go.micro.api.user",err.Error())
	}

	rsp.StatusCode = 200

	getResponse := GetResponse{
		Users: user.MarshalUserCollection(r.Users),
	}

	body, err := json.Marshal(getResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil

}

// GetAll method (User.GetAll) is served by HTTP requests to /user/get-all.
func (e *User) GetAll(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received User.GetAll request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	r, err := e.Client.GetUser(ctx, &protoUser.GetRequest{})
	if err != nil {
		return errors.BadRequest("go.micro.api.user",err.Error())
	}

	rsp.StatusCode = 200

	getResponse := GetResponse{
		Users: user.MarshalUserCollection(r.Users),
	}

	body, err := json.Marshal(getResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Update method (User.Update) is served by HTTP requests to /merchant/user.
func (e *User) Update(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received User.Update request")

	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.user", "This method requires POST")
	}

	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Need Content-Type header")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.example", "Expect application/json")
	}

	var updatedUser *user.User
	err := json.Unmarshal([]byte(req.Body), &updatedUser)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", "Body not valid. Please reference to API documentation.")
	}

	r, err := e.Client.UpdateUser(ctx, user.UnmarshalUser(updatedUser))
	if err != nil {
		return errors.BadRequest("go.micro.api.user",err.Error())
	}

	rsp.StatusCode = 200

	updateResponse := UpdateResponse{
		Updated:   r.Updated,
		User: user.MarshalUser(r.User),
	}

	body, err := json.Marshal(updateResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Delete method (User.Delete) is served by HTTP requests to /user/delete. Full endpoint is /user/delete?id=<user_id>.
func (e *User) Delete(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received User.Delete request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	userID, ok := req.Get["id"]
	if !ok || len(userID.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Please provide an ID")
	}

	r, err := e.Client.DeleteUser(ctx, &protoUser.DeleteRequest{UserID: userID.Values[0]})
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	rsp.StatusCode = 200

	deleteResponse := DeleteResponse{Deleted: r.Deleted}

	body, err := json.Marshal(deleteResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Get method (Transactions.Get) is served by HTTP requests to /user/transaction/get?id=<user_id>.
func (e *Transactions) Get(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received User.Delete request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	userID, ok := req.Get["id"]
	if !ok || len(userID.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Please provide an ID")
	}

	r, err := e.Client.GetTransactions(ctx, &protoTransaction.IDRequest{UserID: userID.Values[0]})
	if err != nil { 
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	rsp.StatusCode = 200

	transactionResponse := TransactionResponse{Transactions: transaction.MarshalTransactionCollection(r.Transactions)}

	body, err := json.Marshal(transactionResponse)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// Set JSON Body
	rsp.Body = string(body)

	return nil
}