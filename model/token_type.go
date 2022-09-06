package model

const (
	GovernanceToken TokenType = iota
	GameToken
	Usdc
	Bnb
	GameVault
	GameNft
)

var tokenTypeShortNames = map[TokenType]string{
	GovernanceToken: "governanceToken",
	GameToken:       "gameToken",
	Bnb:             "bnb",
	GameVault:       "gameVault",
	GameNft:         "gameNft",
}

type TokenType int

func (t TokenType) String() string {
	return tokenTypeShortNames[t]
}
