package constant

const (
	ConfigFile           = "config.toml"
	ConfigEnv            = "ENV_CONFIG"
	PROTOCOL             = "https"
	DOMAIN               = "deep-index.moralis.io"
	MORALIS_API_VERSION  = "api/v2"
	MORALIS_API          = PROTOCOL + "://" + DOMAIN + "/" + MORALIS_API_VERSION + "/"
	NFTTAG               = "nft"
	NFTLISTSUFFIX        = "nft" + "list"
	NFTTYPESUFFIX        = "nft" + "type"
	NATIVETXRECORDSUFFIX = "native" + "txRecord"
	ERC20TXRECORDSUFFIX  = "erc20" + "txRecord"
	NewBlockTopic        = "newBlockTopic"
	BlockConfirmHeight   = 15
	EmptyAddress         = "0x0000000000000000000000000000000000000000"
)

//txType
const (
	GAMETOKEN_RECHARGE = iota + 1
	GOVERNANCETOKEN_RECHARGE
	USDC_RECHARGE
	NATIVE_RECHARGE
	GAMETOKEN_WITHDRAW
	GOVERNANCE_WITHDRAW
	USDC_WITHDRAW
	NATIVE_WITHDRAW
	GAMENFT_TRANSFER
	GAMENFT_IMPORT
	NOT_EXIST
)

//txStatus
const (
	ORDERCREATED = iota
	ORDERHANDLED
	TXFAILED
	TXSUCCESS
)

//tx notify
const (
	NOTNOTIFIED = iota
	NOTIFIED
)

// tx service
const (
	TOKENID      = "tokenId"
	TOKENID_FROM = 100000000
	NONCE        = "nonce"
)
