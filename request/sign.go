package request

type SignRequest struct {
	PrivateKey string `json:"private_key"`
	SignMsg    string `json:"sign_msg"`
}
