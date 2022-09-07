package request

type BatchMintNFTService struct {
	OrderId  string `form:"order_id" json:"order_id" binding:"required"`
	TokenURI string `form:"token_uri" json:"token_uri" binding:"required"`
	Cb       string `form:"cb" json:"cb" binding:"required"`
}
