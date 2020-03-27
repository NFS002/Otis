package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2"
	protoAPI "gitlab.com/otis-team/backend/api/user/proto"
	user "gitlab.com/otis-team/backend/service/user/package"
	protoUser "gitlab.com/otis-team/backend/service/user/proto/user"
	"log"
)

type User struct{
	client protoUser.UserServiceClient
}

type CreatedResponse struct {
	Created bool `json:"created"`
	User *user.User `json:"userID"`
}

type GetResponse struct {
	Users []*user.User `json:"users"`
}

type UpdateResponse struct{
	Updated bool `json:"update"`
	User *user.User `json:"user"`
}

type DeleteResponse struct{
	Deleted bool `json:"deleted"`
}

// User.Create is a method which will be served by http request /user/create
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /user/Create goes to go.micro.api.user User.Create

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

	var newUser *protoUser.User
	err := json.Unmarshal([]byte(req.Body), &newUser)
	if err != nil {
		return errors.BadRequest("go.micro.api.user", "Body not valid. Please reference to API documentation.")
	}

	r, err := e.client.CreateUser(ctx, newUser)
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

	return nil
}

func (e *User) Get(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received User.Get request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Please provide an ID")
	}

	r, err := e.client.GetUser(ctx, &protoUser.GetRequest{UserID: id.Values[0]}) // Seems kinda janky
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

func (e *User) GetAll(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received User.GetAll request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	r, err := e.client.GetUser(ctx, &protoUser.GetRequest{})
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

	r, err := e.client.UpdateUser(ctx, user.UnmarshalUser(updatedUser))
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

func (e *User) Delete(ctx context.Context, req *protoAPI.Request, rsp *protoAPI.Response) error {
	log.Print("Received User.Delete request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.user", "This method requires GET")
	}

	userID, ok := req.Get["id"]
	if !ok || len(userID.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Please provide an ID")
	}

	r, err := e.client.DeleteUser(ctx, &protoUser.DeleteRequest{UserID: userID.Values[0]})
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

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	service.Init()

	client := protoUser.NewUserServiceClient("go.micro.service.user", service.Client())

	m := &User{client}

	// Registering user api handler
	err := protoAPI.RegisterUserHandler(service.Server(), m)
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
