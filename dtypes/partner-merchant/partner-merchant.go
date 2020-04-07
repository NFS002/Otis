package partnermerchant

// PartnerMerchant : Struct is an in-memory representation of a merchant who is part of the Otis platform
// PartnerMerchant resources retrieved from the DB are serialized to an instance of this struct
type PartnerMerchant struct {
	Schema      string `json:"$schema"`
	ID          string `dynamodbav:"id,omitempty" json:"$id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Properties  struct {
		MerchantID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"merchant_id"`
		MerchantName struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"merchant_name"`
		Locations struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Items       struct {
				Description string   `json:"description"`
				Type        []string `json:"type"`
			} `json:"items"`
		} `json:"locations"`
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
		Sector struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"sector"`
		Size struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"size"`
		Rate struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"rate"`
		JoinDate struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"join_date"`
		NextBillingDate struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"next_billing_date"`
		GocardlessMandate struct {
			Description string `json:"description"`
		} `json:"gocardless_mandate"`
		ExpenseBand struct {
			Description string `json:"description"`
			Ref         string `json:"$ref"`
		} `json:"expense_band"`
	} `json:"properties"`
}