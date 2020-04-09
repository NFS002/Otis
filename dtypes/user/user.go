package user

// User : Struct is a local representation of an Otis user
// User resources retrieved from the DB are serialized to an instance of this struct
type User struct {
	ID                   string           		`protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Dob                  string           		`protobuf:"bytes,2,opt,name=Dob,proto3" json:"Dob,omitempty"`
	Gender	 		     int32 			  		`protobuf:"varint,3,opt,name=gender,proto3,enum=user.User_Properties_Gender" json:"gender,omitempty"`
	GenderDescription    string                 `protobuf:"bytes,4,opt,name=genderDescription,proto3" json:"genderDescription,omitempty"`
	University           string                 `protobuf:"bytes,5,opt,name=university,proto3" json:"university,omitempty"`
	JoinDate             string                 `protobuf:"bytes,6,opt,name=joinDate,proto3" json:"joinDate,omitempty"`
	GraduationYear       string                 `protobuf:"bytes,7,opt,name=graduationYear,proto3" json:"graduationYear,omitempty"`
	AverageWeeklySpend   float64                `protobuf:"fixed64,8,opt,name=averageWeeklySpend,proto3" json:"averageWeeklySpend,omitempty"`
	Nationality          string                 `protobuf:"bytes,9,opt,name=nationality,proto3" json:"nationality,omitempty"`
	Tags                 []string               `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
}

// Users represents a slice of User structs
type Users []*User