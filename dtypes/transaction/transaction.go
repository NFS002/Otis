package transaction

// Transaction : Struct is a local representation of a transaction
// Transaction resources retrieved from the DB are serialized to an instance of this struct
type Transaction struct {
	Schema      string `json:"$schema"`
	ID          string `dynamodbav:"id,omitempty" json:"$id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Properties  struct {
		TransactionID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"transaction_id"`
		TransactionTimestamp struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"transaction_timestamp"`
		ProcessedTimestamp struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"processed_timestamp"`
		UserID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"user_id"`
		AccountID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"account_id"`
		MerchantID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"merchant_id"`
		Amount struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"amount"`
		Currency struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"currency"`
		Location struct {
			Description string   `json:"description"`
			Type        []string `json:"type"`
		} `json:"location"`
		Rating struct {
			Description string   `json:"description"`
			Type        []string `json:"type"`
			Minimum     int      `json:"minimum"`
			Maximum     int      `json:"maximum"`
		} `json:"rating"`
		APITags struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Items       struct {
				Type string `json:"type"`
			} `json:"items"`
		} `json:"api_tags"`
		OtisTags struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Items       struct {
				Type string `json:"type"`
			} `json:"items"`
		} `json:"otis_tags"`
		Weather struct {
			Description string `json:"description"`
			Ref         string `json:"$ref"`
		} `json:"weather"`
		Holidays struct {
			Description string `json:"description"`
			Ref         string `json:"$ref"`
		} `json:"holidays"`
		Type struct {
			Description string   `json:"description"`
			Type        string   `json:"type"`
			Enum        []string `json:"enum"`
		} `json:"type"`
		Online struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"online"`
		CashbackValid struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"cashback_valid"`
		MerchantName struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"merchant_name"`
	} `json:"properties"`
	Required []string `json:"required"`
}