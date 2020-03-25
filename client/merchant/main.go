package main

import (
	"encoding/json"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"

	"context"

	pb "gitlab.com/otis-team/backend/client/merchant/proto/merchant"
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
	service := micro.NewService(micro.Name("otis.backend.client.merchant"))
	service.Init()

	client := pb.NewMerchantServiceClient("otis.backend.service.merchant", service.Client())

	//file := defaultFilename
	//if len(os.Args) > 1 {
	//	file = os.Args[1]
	//}

	merchant, err := parseFile("merchant.json")

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateMerchant(context.Background(), merchant)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	//// Retrieve
	//
	//getAll, err := client.GetMerchant(context.Background(), &pb.GetRequest{})
	//if err != nil {
	//	log.Fatalf("Could not list merchants: %v", err)
	//}
	//
	//for _, v := range getAll.Merchants {
	//	log.Println(v)
	//}
}
