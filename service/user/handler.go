package main

import (
	"context"
	"gitlab.com/otis-team/backend/db/client"
	user "gitlab.com/otis-team/backend/dtypes/user/proto"
	userService "gitlab.com/otis-team/backend/service/user/proto/user"
	"log"
)

// Handler struct contains the client connection to the DB, to be used by Handler functions.
type Handler struct {
	Client client.DynamoClient
}

// CreateUser handles gRPC requests to create a new user in the DB.
func (h *Handler) CreateUser(ctx context.Context, req *user.User, res *userService.UsersResponse) error {
	log.Print("CreateUser handler fired")
	_, err := h.Client.CreateUser(req)
	res.Executed = err == nil
	res.Users = []*user.User{ req }
	return nil
}

// GetUser handles gRPC requests to retrieve one (if User ID is supplied) or many users from the DB.
func (h *Handler) GetUser(ctx context.Context, req *userService.UserQuery, res *userService.UsersResponse) error {
	log.Print("GetUser handler fired")

	var err error
	var dbUsers []*user.User
	var dbUser *user.User

	if len(req.UserID) == 0 {
		dbUsers, err = h.Client.GetAllUsers()
		res.Users = dbUsers
		res.Executed = err == nil
	} else {
		dbUser, err = h.Client.GetUserByID(req.UserID)
		res.Users = []*user.User{ dbUser }
		res.Executed = err == nil
	}

	return err
}

// DeleteUser handles gRPC requests to delete a new user from the DB
func (h *Handler) DeleteUser(ctx context.Context, req *userService.UserQuery, res *userService.UsersResponse) error {
	log.Print("DeleteUser handler fired!")
	err := h.Client.DeleteUser(req.UserID)
	res.Executed = err == nil
	return err
}