package constant

const (
	ConfigFile          = "config.toml"
	ConfigEnv           = "ENV_CONFIG"
	PROTOCOL            = "https"
	DOMAIN              = "deep-index.moralis.io"
	MORALIS_API_VERSION = "api/v2"
	MORALIS_API         = PROTOCOL + "://" + DOMAIN + "/" + MORALIS_API_VERSION + "/"
)
