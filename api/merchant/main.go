package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	proto "gitlab.com/otis-team/backend/api/merchant/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
)

type Merchant struct{}

// Merchant.Create is a method which will be served by http request /merchant/create
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /merchant/Create goes to go.micro.api.merchant Merchant.Create

func (e *Merchant) Create(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Example.Call request")

	// parse values from the get request
	id, ok := req.Get["id"]

	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.api.merchant", "no content")
	}

	// set response status
	rsp.StatusCode = 200

	// Need to add POST functionality and push to DB

	// respond with some json
	b, _ := json.Marshal(map[string]string{
		"message": "Will create user with ID:  " + strings.Join(id.Values, " "),
	})

	// set json body
	rsp.Body = string(b)

	return nil
}

func (e *Merchant) Get(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Merchant.Get request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.merchant", "Please provide an ID")
	}

	// Need to retrieve user heree

	body, _ := json.Marshal(map[string]string{
		"message": "We'll get the user with ID: " + strings.Join(id.Values, " "),
	})

	rsp.StatusCode = 200
	rsp.Body = string(body)

	return nil
}

func (e *Merchant) GetAll(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Merchant.GetAll request")

	if req.Method != "GET" {
		return errors.BadRequest("go.micro.api.merchant", "This method requires GET")
	}

	// Need to retrieve all users here

	body, _ := json.Marshal(map[string]string{
		"message": "We'll get all users",
	})

	rsp.StatusCode = 200
	rsp.Body = string(body)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.merchant"),
	)

	service.Init()

	// Registering merchant handler
	proto.RegisterMerchantHandler(service.Server(), new(Merchant))


	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
