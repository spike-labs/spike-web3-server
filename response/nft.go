package response

type Metadata struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	ExternalUrl string    `json:"external_url"`
	Fid         string    `json:"fid"`
	Mid         string    `json:"mid"`
	SpikeInfo   SpikeInfo `json:"spike_info"`
	Attribute   []Attr    `json:"attributes"`
}

type SpikeInfo struct {
	Version string `json:"version"`
	Tp      string `json:"type"`
	Url     string `json:"url"`
}

type Attr struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
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
