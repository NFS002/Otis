package merchant

import (
	"context"
	"log"
	"os"
	"fmt"

	pb "gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	"github.com/micro/go-micro"

)

const (
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(
			micro.Name("go.micro.service.merchant"),
		)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	merchantCollection := client.Database("otis").Collection("merchant")

	repository := &MongoRepository{merchantCollection}
	h := &handler{repository}

	pb.RegisterMerchantServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}