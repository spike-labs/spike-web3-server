package model

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

type CacheData struct {
	Type        string                 `json:"type"`
	GameId      string                 `json:"gameId"`
	BlockNumber string                 `json:"block_number"`
	TokenId     string                 `json:"token_id"`
	Image       string                 `json:"image"`
	Description string                 `json:"description"`
	SpikeInfo   SpikeInfo              `json:"spike_info"`
	Attributes  map[string]interface{} `json:"attributes"`
}
