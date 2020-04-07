package transaction

// Holiday : Struct is an in-memory representation of a time of year, such as a holiday
// Holiday resources retrieved from the DB are serialized to an instance of this struct
type Holiday struct {
	Schema      string `json:"$schema"`
	ID          string `dynamodbav:"id,omitempty" json:"$id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Properties  struct {
		Christmas struct {
			Type string `json:"type"`
		} `json:"christmas"`
		Easter struct {
			Type string `json:"type"`
		} `json:"easter"`
		BankHoliday struct {
			Type string `json:"type"`
		} `json:"bank_holiday"`
	} `json:"properties"`
}