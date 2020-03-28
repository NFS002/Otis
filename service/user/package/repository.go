package user

import (
	"context"
	"github.com/satori/go.uuid"
	pb "gitlab.com/otis-team/backend/service/user/proto/user"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User struct maps protobuf definition. Contains json and bson key mappings.
type User struct {
	UserID string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DoB string `json:"dob"`
	Gender string `json:"gender"`
	University string `json:"university"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Accounts Accounts `json:"accounts"`
}

// Account struct maps protobuf definition. Contains json and bson key mappings.
type Account struct {
	AccountID string `json:"account_id"`
}

// Accounts struct represents slice of Account structs
type Accounts []*Account

// Accounts

// MarshalAccountCollection converts slice of Account protobufs to slice of Account structs
func MarshalAccountCollection(accounts []*pb.Account) []*Account {
	Collection := make([]*Account, 0)
	for _, account := range accounts {
		Collection = append(Collection, MarshalAccount(account))
	}
	return Collection
}

// UnmarshalAccountCollection converts slice of Account structs to slice of Account protobufs
func UnmarshalAccountCollection(accounts []*Account) []*pb.Account {
	Collection := make([]*pb.Account, 0)
	for _, account := range accounts {
		Collection = append(Collection, UnmarshalAccount(account))
	}
	return Collection
}

// MarshalAccount converts Account protobuf to Account struct
func MarshalAccount(account *pb.Account) *Account {
	return &Account{
		AccountID: account.AccountID,
	}
}

// UnmarshalAccount converts Account struct to Account protobuf
func UnmarshalAccount(account *Account) *pb.Account {
	return &pb.Account{
		AccountID: account.AccountID,
	}
}

// User

// MarshalUserCollection converts slice of User protobufs to slice of User structs
func MarshalUserCollection(users []*pb.User) []*User {
	Collection := make([]*User, 0)
	for _, user := range users {
		Collection = append(Collection, MarshalUser(user))
	}
	return Collection
}

// UnmarshalUserCollection converts slice of User structs to slice of User protobufs
func UnmarshalUserCollection(users []*User) []*pb.User {
	Collection := make([]*pb.User, 0)
	for _, user := range users {
		Collection = append(Collection, UnmarshalUser(user))
	}
	return Collection
}

// MarshalUser converts User protobuf to User struct
func MarshalUser(user *pb.User) *User {
	accounts := MarshalAccountCollection(user.Accounts)

	return &User{
		UserID:     user.UserID,
		FirstName:  user.FirstName,
		LastName:  	user.LastName,
		DoB:        user.Dob,
		Gender:     user.Gender,
		University: user.University,
		Email:      user.Email,
		Phone:      user.Phone,
		Accounts:   accounts,
	}
}

// UnmarshalUser converts User struct to User protobuf
func UnmarshalUser(user *User) *pb.User {
	accounts := UnmarshalAccountCollection(user.Accounts)

	return &pb.User{
		UserID:     user.UserID,
		FirstName:  user.FirstName,
		LastName:  	user.LastName,
		Dob:        user.DoB,
		Gender:     user.Gender,
		University: user.University,
		Email:      user.Email,
		Phone:      user.Phone,
		Accounts:   accounts,
	}
}

// Repository

// Repository interface describes all available repository methods. Currently basic CRUD.
type Repository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, userID string) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID string) error
}

// MongoRepository struct describes specific collection relevant to the repository being used.
type MongoRepository struct {
	Collection *mongo.Collection
}

// Create method implements functionality to create a user in the DB. UUID is generated for user_id.
func (repository *MongoRepository) Create(ctx context.Context, user *User) (*User, error) {
	uuid, err := generateUUID()
	if err != nil {
		return nil, err
	}

	user.UserID = uuid.String()

	_, err = repository.Collection.InsertOne(ctx, user)

	return user, err
}

// GetAll method implements functionality to retrieve all users from the DB.
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*User, error) {
	cur, err := repository.Collection.Find(ctx, bson.D{}, nil)
	var users []*User
	for cur.Next(ctx) {
		var user *User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}

// Get method implements functionality to retrieve users from teh DB matching userID.
//Possible for multiple to be returned.
func (repository *MongoRepository) Get(ctx context.Context, userID string) ([]*User, error) {
	cur, err := repository.Collection.Find(ctx, bson.M{"user_id": userID}, nil)
	var users []*User
	for cur.Next(ctx) {
		var user *User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}

// Update method implements functionality to update a user in the DB matching user ID.
// Only requirement for supplied new user is that the user field matches an existing
// user in the DB.
func (repository *MongoRepository) Update(ctx context.Context, user *User) error {
	err := repository.Collection.FindOneAndReplace(ctx, bson.M{"user_id": user.UserID}, user)
	if err != nil {
		return err.Err()
	}

	return nil
}

// Delete method implements functionality to delete a user from the DB matching user ID.
func (repository *MongoRepository) Delete(ctx context.Context, userID string) error {
	err := repository.Collection.FindOneAndDelete(ctx, bson.M{"user_id": userID})
	if err != nil {
		return err.Err()
	}

	return nil
}

// UUID

// generateUUID generates a random UUID.
func generateUUID() (uuid.UUID, error) {
	var err error

	uuid := uuid.Must(uuid.NewV4(), err)

	return uuid, err
}










