package response

type NftType struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type NftResults struct {
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Cursor   string      `json:"cursor"`
	Results  []NftResult `json:"result"`
}

type NftResult struct {
	TokenId     string `json:"token_id"`
	BlockNumber string `json:"block_number"`
	TokenUri    string `json:"token_uri"`
	Metadata    string `json:"metadata"`
}
