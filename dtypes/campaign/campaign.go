package campaign

// Campaign : Struct is an in-memory representation of a campaign resource
// Campaign resources retrieved fron the DB are serialized to an instance of this struct
type Campaign struct {
	Schema      string `json:"$schema"`
	ID          string `dynamodbav:"id,omitempty" json:"$id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Properties  struct {
		CampaignID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"campaign_id"`
		MerchantID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"merchant_id"`
		CampaignRate struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"campaign_rate"`
		BeginTimestamp struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"begin_timestamp"`
		EndTimestamp struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"end_timestamp"`
		UserSet struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Items       struct {
				Type string `json:"type"`
				Ref  string `json:"$ref"`
			} `json:"items"`
		} `json:"user_set"`
		TransactionsSet struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Items       struct {
				Type string `json:"type"`
				Ref  string `json:"$ref"`
			} `json:"items"`
		} `json:"transactions_set"`
		ValidFor struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Minimum     int    `json:"minimum"`
			Default     int    `json:"default"`
		} `json:"valid_for"`
		CampaignTitle struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"campaign_title"`
		CampaignDescription struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"campaign_description"`
		Locations struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Items       struct {
				Type string `json:"type"`
			} `json:"items"`
		} `json:"locations"`
		Active struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"active"`
	} `json:"properties"`
}