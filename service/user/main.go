package main

import (
	"context"
	"fmt"
	user "gitlab.com/otis-team/backend/service/user/package"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "gitlab.com/otis-team/backend/service/user/proto/user"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := user.CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	userCollection := client.Database("otis").Collection("user")

	repository := &user.MongoRepository{userCollection}
	h := &user.Handler{repository}

	pb.RegisterUserServiceHandler(service.Server(), h)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}