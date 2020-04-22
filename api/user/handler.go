package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/errors"
	proto "gitlab.com/otis_team/backend/api/user/proto"
	user "gitlab.com/otis_team/backend/dtypes/user/proto"
	transactionService "gitlab.com/otis_team/backend/service/transaction/proto/transaction"
	userService "gitlab.com/otis_team/backend/service/user/proto/user"
	"log"
)

// User struct. All methods using this struct will be mapped to /user/<method>.
type User struct {
	Client userService.UserService
}

// Transactions struct. All methods using this struct will be mapped to /user/transaction/<method>
type Transactions struct {
	Client transactionService.TransactionService
}

// Create method (User.Create) is served by HTTP requests to /user/create.
func (e *User) Create(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
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

	newUser := &user.User{}
	err := json.Unmarshal([]byte(req.Body), newUser)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", "Body not valid. Please reference to API documentation.")
	}

	r, err := e.Client.CreateUser(ctx, newUser)
	if err != nil {
		return errors.BadRequest("go.micro.api.user",err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"executed":   r.Executed,
		"user": r.Users,
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Get method (User.Get) is served by HTTP requests to /user/get. Full endpoint is /user/get?id=<user_id>.
func (e *User) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received User.Get request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Please provide an ID")
	}

	r, err := e.Client.GetUser(ctx, &userService.UserQuery{UserID: id.Values[0]}) // Seems kinda janky
	if err != nil {
		return errors.BadRequest("go.micro.api.user",err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"users": r.Users,
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil

}

// GetAll method (User.GetAll) is served by HTTP requests to /user/get-all.
func (e *User) GetAll(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received User.GetAll request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	r, err := e.Client.GetUser(ctx, &userService.UserQuery{})
	if err != nil {
		return errors.BadRequest("go.micro.api.user",err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"users": r.Users,
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Update method (User.Update) is served by HTTP requests to /merchant/user.
func (e *User) Update(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
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

	updatedUser := &user.User{}
	err := json.Unmarshal([]byte(req.Body), &updatedUser)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", "Body not valid. Please reference to API documentation.")
	}

	/* UpdateUser calls CreateUser internally */
	r, err := e.Client.CreateUser(ctx, updatedUser)
	if err != nil {
		return errors.BadRequest("go.micro.api.user",err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"executed":   r.Executed,
		"users": r.Users,
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Delete method (User.Delete) is served by HTTP requests to /user/delete. Full endpoint is /user/delete?id=<user_id>.
func (e *User) Delete(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received User.Delete request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	userID, ok := req.Get["id"]
	if !ok || len(userID.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Please provide an ID")
	}

	r, err := e.Client.DeleteUser(ctx, &userService.UserQuery{UserID: userID.Values[0]})
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"executed": r.Executed,
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	// set json body
	rsp.Body = string(body)

	return nil
}

// Get method (Transactions.Get) is served by HTTP requests to /user/transaction/get?id=<user_id>.
func (e *Transactions) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received User.Delete request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	userID, ok := req.Get["id"]
	if !ok || len(userID.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Please provide an ID")
	}

	r, err := e.Client.GetTransactions(ctx, &transactionService.TransactionQuery{UserID: userID.Values[0]})
	if err != nil { 
		return errors.BadRequest("go.micro.api.user", err.Error())
	}

	rsp.StatusCode = 200

	responseBody := map[string]interface{}{
		"transactions": r.Transactions,
	}

	body, err := json.Marshal(responseBody)
	if err != nil {
		return errors.BadRequest("go.micro.api.merchant", err.Error())
	}

	// Set JSON Body
	rsp.Body = string(body)

	return nil
}
