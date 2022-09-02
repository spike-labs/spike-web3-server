package request

type BatchMintNFTService struct {
	TokenID  int64  `form:"token_id" json:"token_id" binding:"required"`
	TokenURI string `form:"token_uri" json:"token_uri" binding:"required"`
}
