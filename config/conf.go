package config

import (
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("config")

var Cfg Config

type Config struct {
	System         System         `json:"system" toml:"System"`
	Moralis        Moralis        `json:"moralis" toml:"Moralis"`
	BscScan        BscScan        `json:"bscscan" toml:"Bscscan"`
	Contract       Contract       `json:"contract" toml:"Contract"`
	Redis          Redis          `json:"redis" toml:"Redis"`
	Chain          Chain          `json:"chain" toml:"Chain"`
	Mysql          Mysql          `json:"mysql" toml:"Mysql"`
	SignService    SignService    `json:"signService" toml:"SignService"`
	SignWorkers    []SignWorker   `json:"signWorkers" toml:"SignWorkers"`
	Limit          Limit          `json:"limit" toml:"Limit"`
	TxApiWhiteList TxApiWhiteList `toml:"TxApiWhiteList"`
}

type TxApiWhiteList struct {
	IpList []string `toml:"Ip"`
}

type Limit struct {
	NftLimit      int `json:"nftLimit" toml:"NftLimit"`
	TxRecordLimit int `json:"txRecordLimit" toml:"TxRecordLimit"`
}

type Mysql struct {
	Path         string `json:"path" toml:"Path"`
	Port         string `json:"port" toml:"Port"`
	Config       string `json:"config" toml:"Config"`
	Dbname       string `json:"db_name" toml:"DbName"`
	Username     string `json:"username" toml:"Username"`
	Password     string `json:"password" toml:"Password"`
	MaxIdleConns int    `json:"maxIdleConns" toml:"MaxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" toml:"MaxOpenConns"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

type Chain struct {
	WsNodeAddress  string `toml:"WsNodeAddress"`
	RpcNodeAddress string `toml:"RpcNodeAddress"`
}

type System struct {
	Port          string `toml:"Port"`
	MachineId     string `toml:"MachineId"`
	SignSecretKey string `toml:"SignSecretKey"`
}

type Moralis struct {
	XApiKey string `toml:"XApiKey"`
}

type BscScan struct {
	ApiKey    string `toml:"ApiKey"`
	UrlPrefix string `toml:"UrlPrefix"`
}

type Redis struct {
	Address  string `toml:"Address"`
	Password string `toml:"Password"`
}

type Contract struct {
	NftContractAddress   []string `toml:"NftContract"`
	ERC20ContractAddress []string `toml:"ERC20Contract"`
	GameVaultAddress     string   `toml:"GameVaultAddress"`
}

type SignService struct {
	TaskThreshold int `toml:"TaskThreshold"`
	SchedInterval int `toml:"SchedInterval"`
}

type SignWorker struct {
	WalletAddress string `toml:"WalletAddress"`
	ServerUrl     string `toml:"ServerUrl"`
}
