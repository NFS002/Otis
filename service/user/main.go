package main

import (
	"context"
	"sync"
	"fmt"

	pb "gitlab.com/otis-team/backend/service/user/proto/user"
	"github.com/micro/go-micro"

)

const (
	port = ":50051"
)

type repository interface {
	Create(user *pb.User) (*pb.User, error)
	GetAll() []*pb.User
	Get(id string) ([]*pb.User, error)
}

type Repository struct {
	mu sync.RWMutex
	users []*pb.User
}

func (repo *Repository) Create(user *pb.User) (*pb.User, error) {
	repo.mu.Lock()
	updated := append(repo.users, user)
	repo.users = updated
	repo.mu.Unlock()
	return user, nil
}

func (repo *Repository) GetAll() []*pb.User {
	return repo.users
}

func (repo *Repository) Get(id string) ([]*pb.User, error) {
	var res []*pb.User

	for _, user := range repo.users {
		if user.UserID == id {
			res = append(res, user)
		}
	}

	return res, nil
}

type Service struct {
	repo repository
}

func (s *service) CreateUser(ctx context.Context, req *pb.User, res *pb.CreateResponse) error {
	user, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.User = user
	res.User.UserID = "1234"

	return nil
}

// REFACTOR
func (s *service) GetUser(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	if req.Id != "" {
		users, err := s.repo.Get(req.Id)
		if err != nil {
			return err
		}
		res.Users = users
	} else {
		users := s.repo.GetAll()

		res.Users = users
	}

	return nil
}

func main() {
	repo := &Repository{}

	service := micro.NewService(
		micro.Name("otis-app.com.user.service"),
		)

	service.Init()

	pb.RegisterUserServiceHandler(service.Server(), &Service{repo})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}