package config

import (
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("config")

var Cfg Config

type Config struct {
	System      System       `json:"system" toml:"system"`
	Moralis     Moralis      `json:"moralis" toml:"moralis"`
	BscScan     BscScan      `json:"bscscan" toml:"bscscan"`
	Contract    Contract     `json:"contract" toml:"contract"`
	Redis       Redis        `json:"redis" toml:"redis"`
	Chain       Chain        `json:"chain" toml:"chain"`
	Mysql       Mysql        `json:"mysql" toml:"mysql"`
	SignService SignService  `json:"signService" toml:"signService"`
	SignWorkers []SignWorker `json:"signWorkers" toml:"signWorkers"`
	Model       Model        `json:"model" toml:"model"`
	Limit       Limit        `json:"limit" toml:"limit"`
}

type Model struct {
	Name []string `json:"name" toml:"name"`
}

type Limit struct {
	NftLimit      int `json:"nftLimit" toml:"nftLimit"`
	TxRecordLimit int `json:"txRecordLimit" toml:"txRecordLimit"`
}

type Mysql struct {
	Path         string `json:"path" toml:"path"`
	Port         string `json:"port" toml:"port"`
	Config       string `json:"config" toml:"config"`
	Dbname       string `json:"db_name" toml:"dbName"`
	Username     string `json:"username" toml:"username"`
	Password     string `json:"password" toml:"password"`
	MaxIdleConns int    `json:"maxIdleConns" toml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" toml:"maxOpenConns"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

type Chain struct {
	WsNodeAddress  string `toml:"wsNodeAddress"`
	RpcNodeAddress string `toml:"rpcNodeAddress"`
}

type System struct {
	Port      string `toml:"port"`
	MachineId string `toml:"machineId"`
}

type Moralis struct {
	XApiKey string `toml:"xApiKey"`
}

type BscScan struct {
	ApiKey    string `toml:"apiKey"`
	UrlPrefix string `toml:"urlPrefix"`
}

type Redis struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
}

type Contract struct {
	GameNftAddress         string `toml:"gameNftAddress"`
	GovernanceTokenAddress string `toml:"governanceTokenAddress"`
	GameTokenAddress       string `toml:"gameTokenAddress"`
	GameVaultAddress       string `toml:"gameVaultAddress"`
	UsdcAddress            string `toml:"usdcAddress"`
}

type SignService struct {
	TaskThreshold int `toml:"taskThreshold"`
	SchedInterval int `toml:"schedInterval"`
}

type SignWorker struct {
	WalletAddress string `toml:"walletAddress"`
	ServerUrl     string `toml:"serverUrl"`
}
