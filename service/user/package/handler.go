package user

import (
	"context"
	pb "gitlab.com/otis-team/backend/service/user/proto/user"
	"log"
)

// Handler struct contains other structs (mainly Repository, aka the client connection to the DB) to be used by Handler functions.
type Handler struct {
	Repository
}

// CreateUser handles gRPC requests to create a new user in the DB.
func (s *Handler) CreateUser(ctx context.Context, req *pb.User, res *pb.CreateResponse) error {
	log.Print("CreateUser handler fired!")

	user, err := s.Repository.Create(ctx, MarshalUser(req))
	if err != nil {
		res.Created = true
		res.User = UnmarshalUser(user)
	}

	return nil
}

// GetUser handles gRPC requests to retrieve one (if User ID is upplied) or many users from the DB.
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

// UpdateUser handles gRPC requests to update a new user in the DB
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

// DeleteUser handles gRPC requests to delete a new user from the DB
func (s *Handler) DeleteUser(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	log.Print("DeleteUser handler fired!")

	err := s.Repository.Delete(ctx, req.UserID)
	if err != nil {
		return err
	}

	res.Deleted = true
	return nil
}