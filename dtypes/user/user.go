package user

import pb "gitlab.com/otis-team/backend/service/user/proto/user"

// User : Struct is a local representation of an Otis user
// User resources retrieved from the DB are serialized to an instance of this struct
type User struct {
	Schema      string `json:"$schema"`
	ID          string `dynamodbav:"id,omitempty" json:"$id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Properties  struct {
		UserID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"user_id"`
		DateOfBirth struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"date_of_birth"`
		Gender struct {
			Description string   `json:"description"`
			Type        string   `json:"type"`
			Enum        []string `json:"enum"`
		} `json:"gender"`
		SpecificGender struct {
			Description string   `json:"description"`
			Type        []string `json:"type"`
		} `json:"specific_gender"`
		University struct {
			Description string `json:"description"`
			Ref         string `json:"$ref"`
		} `json:"university"`
		JoinDate struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"join_date"`
		GraduationYear struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"graduation_year"`
		AverageWeeklySpend struct {
			Description string   `json:"description"`
			Type        []string `json:"type"`
		} `json:"average_weekly_spend"`
		Nationality struct {
			Description string   `json:"description"`
			Type        []string `json:"type"`
		} `json:"nationality"`
		Tags struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Items       struct {
				Type []string `json:"type"`
			} `json:"items"`
		} `json:"tags"`
		ExpenseBands struct {
			Description string `json:"description"`
			Ref         string `json:"$ref"`
		} `json:"expense_bands"`
	} `json:"properties"`
}