package response

type BalanceShow struct {
	Symbol  string `json:"symbol"`
	Balance string `json:"balance"`
}
