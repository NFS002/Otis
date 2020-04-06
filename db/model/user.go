package model

import pb "gitlab.com/otis-team/backend/service/user/proto/user"

// User : struct to represent a DynamoDb user resource
// Protobuf user messages must be marshalled into a instance of this struct
// before then being marhalled to a dynamodb attribute value.
// Similarly, users retrieved from dynamodb will be represented as an instance of this struct.
type User struct {
	UserID string `json:"user_id" dynamodbav:"user_id"`
	FirstName string `json:"first_name" dynamodbav:"first_name"`
	LastName string `json:"last_name" dynamodbav:"last_name"`
	DoB string `json:"dob" dynamodbav:"dob"`
	Gender string `json:"gender" dynamodbav:"gender"`
	University string `json:"university" dynamodbav:"university"`
	Email string `json:"email" dynamodbav:"email"`
	Phone string `json:"phone" dynamodbav:"phone"`
	Accounts Accounts `json:"accounts" dynamodbav:"accounts"`
}

// Users struct represents a slice of User structs
type Users []*User

// Account struct maps protobuf definition.
type Account struct {
	AccountID string `json:"account_id" dynamodbav:"account_id"`
}

// Accounts struct represents slice of Account structs
type Accounts []*Account

// ProtobufToUser : Converts a User protobuf message to User struct
func ProtobufToUser(user *pb.User) *User {
	accounts := ProtobufToAccountCollection(user.Accounts)
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

// UserToProtobuf : Converts User struct to User protobuf message
func UserToProtobuf(user *User) *pb.User {

	accounts := AccountCollectionToProtobuf(user.Accounts)

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

// ProtobufToUserCollection : Converts slice of User protobuf messages to slice of User structs
func ProtobufToUserCollection(users []*pb.User) []*User {
	Collection := make([]*User, 0)
	for _, user := range users {
		Collection = append(Collection, ProtobufToUser(user))
	}
	return Collection
}

// UserCollectionToProtobuf : Converts slice of User structs to slice of User protobuf messages
func UserCollectionToProtobuf(users []*User) []*pb.User {
	Collection := make([]*pb.User, 0)
	for _, user := range users {
		Collection = append(Collection, UserToProtobuf(user))
	}
	return Collection
}

// ProtobufToAccountCollection : Converts slice of Account protobuf messages to slice of Account structs
func ProtobufToAccountCollection(accounts []*pb.Account) []*Account {
	Collection := make([]*Account, 0)
	for _, account := range accounts {
		Collection = append(Collection, ProtobufToAccount(account))
	}
	return Collection
}

// AccountCollectionToProtobuf : Converts slice of Account structs to slice of Account protobuf messages
func AccountCollectionToProtobuf(accounts []*Account) []*pb.Account {
	Collection := make([]*pb.Account, 0)
	for _, account := range accounts {
		Collection = append(Collection, AccountToProtobuf(account))
	}
	return Collection
}

// ProtobufToAccount : Converts Account protobuf message to an Account struct
func ProtobufToAccount(account *pb.Account) *Account {
	return &Account{
		AccountID: account.AccountID,
	}
}

// AccountToProtobuf : Converts Account struct to Account protobuf message
func AccountToProtobuf(account *Account) *pb.Account {
	return &pb.Account{
		AccountID: account.AccountID,
	}
}