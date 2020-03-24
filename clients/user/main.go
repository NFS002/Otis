package main

import (
"encoding/json"
"github.com/micro/go-micro"
"io/ioutil"
"log"

"context"

pb "gitlab.com/otis-team/backend/clients/user/proto/user"
)

const (
	address = "localhost:50051"
	defaultFilename = "merchant.json"
)

func parseFile(file string) (*pb.User, error) {
	var user *pb.User
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &user)

	return user, err
}

func main() {
	service := micro.NewService(micro.Name("otis-app.com.user.client"))
	service.Init()

	client := pb.NewUserServiceClient("otis-app.com.user.service", service.Client())

	//file := defaultFilename
	//if len(os.Args) > 1 {
	//	file = os.Args[1]
	//}

	merchant, err := parseFile("/app/user.json")

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateUser(context.Background(), merchant)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	// Retrieve

	getAll, err := client.GetUser(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}

	for _, v := range getAll.Users {
		log.Println(v)
	}
}
