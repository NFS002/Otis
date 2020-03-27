package user

import (
	"context"
	pb "gitlab.com/otis-team/backend/service/user/proto/user"
	"log"
)

type Handler struct {
	Repository
}

func (s *Handler) CreateUser(ctx context.Context, req *pb.User, res *pb.CreateResponse) error {
	log.Print("CreateUser handler fired!")

	user, err := s.Repository.Create(ctx, MarshalUser(req))
	if err != nil {
		res.Created = true
		res.User = UnmarshalUser(user)
	}

	return nil
}

func (s *Handler) GetUser(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	log.Print("GetUser handler fired!")

	var users []*User
	var err error

	if len(req.UserID) == 0 {
		users, err = s.Repository.GetAll(ctx)
	} else {
		users, err = s.Repository.Get(ctx, req.UserID)
	}

	res.Users = UnmarshalUserCollection(users)

	return err
}

func (s *Handler) UpdateUser(ctx context.Context, req *pb.User, res *pb.UpdateResponse) error {
	log.Print("UpdateUser handler fired!")

	err := s.Repository.Update(ctx, MarshalUser(req))
	if err != nil {
		return err
	}

	res.Updated = true
	res.User = req

	return nil
}

func (s *Handler) DeleteUser(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	log.Print("DeleteUser handler fired!")

	err := s.Repository.Delete(ctx, req.UserID)
	if err != nil {
		return err
	}

	res.Deleted = true
	return nil
}