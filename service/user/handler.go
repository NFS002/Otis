package main

import (
	"context"
	"gitlab.com/otis-team/backend/db/client"
	"gitlab.com/otis-team/backend/db/model"
	pb "gitlab.com/otis-team/backend/service/user/proto/user"
	"log"
)

// Handler struct contains the client connection to the DB, to be used by Handler functions.
type Handler struct {
	Client client.DynamoClient
}

// CreateUser handles gRPC requests to create a new user in the DB.
func (h *Handler) CreateUser(ctx context.Context, req *pb.User, res *pb.CreateResponse) error {
	log.Print("CreateUser handler fired")
	user := model.ProtobufToUser(req)
	_, err := h.Client.CreateUser(user)
	if err != nil {
		return err
	}
	res.Created = true
	res.User = req
	return nil
}

// GetUser handles gRPC requests to retrieve one (if User ID is supplied) or many users from the DB.
func (h *Handler) GetUser(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	log.Print("GetUser handler fired")

	var err error
	var users *model.Users

	if len(req.UserID) == 0 {
		users, err = h.Client.GetAllUsers()
	} else {
		users, err = h.Client.GetUserByID(req.UserID)
	}
	if err != nil {
		log.Println(err.Error())
		return err
	}

	res.Users = model.UserCollectionToProtobuf(*users)
	return nil
}

// UpdateUser handles gRPC requests to update a new user in the DB
func (h *Handler) UpdateUser(ctx context.Context, req *pb.User, res *pb.UpdateResponse) error {
	log.Print("UpdateUser handler fired")
	user := model.ProtobufToUser( req )
	_, err := h.Client.CreateUser(user)
	if err != nil {
		return err
	}
	res.Updated = true
	res.User = req
	return nil
}

// DeleteUser handles gRPC requests to delete a new user from the DB
func (h *Handler) DeleteUser(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	log.Print("DeleteUser handler fired!")
	err := h.Client.DeleteUser(req.UserID)
	res.Deleted = (err == nil)
	return err
}