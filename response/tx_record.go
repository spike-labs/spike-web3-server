package response

type TxResult struct {
	Hash        string `json:"hash"`
	TimeStamp   string `json:"timeStamp"`
	BlockNumber string `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Input       string `json:"input"`
	Type        string `json:"type"`
}

type BscResult struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Result  []TxResult `json:"result"`
}
