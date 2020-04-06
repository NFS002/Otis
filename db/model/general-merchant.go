package model

// GeneralMerchant : Struct is an in-memory representation of a merchant who is not part of the Otis platform
// GeneralMerchant resources retrieved from the DB are serialized to an instance of this struct
type GeneralMerchant struct {
	Schema      string `json:"$schema"`
	ID          string `json:"$id"`
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
		ScrapeTags struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Items       struct {
				Type string `json:"type"`
			} `json:"items"`
		} `json:"scrape_tags"`
		Sector struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"sector"`
		Size struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"size"`
		ExpenseBand struct {
			Description string `json:"description"`
			Ref         string `json:"$ref"`
		} `json:"expense_band"`
	} `json:"properties"`
}