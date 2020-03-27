package user

import (
	"context"
	"github.com/satori/go.uuid"
	pb "gitlab.com/otis-team/backend/service/user/proto/user"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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

type Account struct {
	AccountID string `json:"account_id"`
}

type Accounts []*Account

// Accounts

func MarshalAccountCollection(accounts []*pb.Account) []*Account {
	Collection := make([]*Account, 0)
	for _, account := range accounts {
		Collection = append(Collection, MarshalAccount(account))
	}
	return Collection
}

func UnmarshalAccountCollection(accounts []*Account) []*pb.Account {
	Collection := make([]*pb.Account, 0)
	for _, account := range accounts {
		Collection = append(Collection, UnmarshalAccount(account))
	}
	return Collection
}

func MarshalAccount(account *pb.Account) *Account {
	return &Account{
		AccountID: account.AccountID,
	}
}

func UnmarshalAccount(account *Account) *pb.Account {
	return &pb.Account{
		AccountID: account.AccountID,
	}
}

// User

func MarshalUserCollection(users []*pb.User) []*User {
	Collection := make([]*User, 0)
	for _, user := range users {
		Collection = append(Collection, MarshalUser(user))
	}
	return Collection
}

func UnmarshalUserCollection(users []*User) []*pb.User {
	Collection := make([]*pb.User, 0)
	for _, user := range users {
		Collection = append(Collection, UnmarshalUser(user))
	}
	return Collection
}

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

type Repository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, userID string) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID string) error
}

type MongoRepository struct {
	Collection *mongo.Collection
}

func (repository *MongoRepository) Create(ctx context.Context, user *User) (*User, error) {
	uuid, err := generateUUID()
	if err != nil {
		return nil, err
	}

	user.UserID = uuid.String()

	_, err = repository.Collection.InsertOne(ctx, user)

	return user, err
}

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

func (repository *MongoRepository) Update(ctx context.Context, user *User) error {
	err := repository.Collection.FindOneAndReplace(ctx, bson.M{"user_id": user.UserID}, user)
	if err != nil {
		return err.Err()
	}

	return nil
}

func (repository *MongoRepository) Delete(ctx context.Context, userID string) error {
	err := repository.Collection.FindOneAndDelete(ctx, bson.M{"user_id": userID})
	if err != nil {
		return err.Err()
	}

	return nil
}

// UUID

func generateUUID() (uuid.UUID, error) {
	var err error

	uuid := uuid.Must(uuid.NewV4(), err)

	return uuid, err
}










