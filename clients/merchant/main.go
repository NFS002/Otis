package main

import (
	"encoding/json"
	micro "github.com/micro/go-micro"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "otis-app.com/backend/clients/merchant/proto/merchant"

)

const (
	address = "localhost:50051"
	defaultFilename = "merchant.json"
)

func parseFile(file string) (*pb.Merchant, error) {
	var merchant *pb.Merchant
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &merchant)

	return merchant, err
}

func main() {
	service := micro.NewService(micro.Name("otis-app.com.client.merchant"))
	service.Init()

	client := pb.NewMerchantServiceClient("otis-app.com.service.merchant", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	merchant, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateMerchant(context.Background(), merchant)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	// Retrieve

	getAll, err := client.GetMerchant(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list merchants: %v", err)
	}

	for _, v := range getAll.Merchants {
		log.Println(v)
	}
}
