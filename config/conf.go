package config

import (
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("config")

var Cfg Config

type Config struct {
	System      System       `toml:"system"`
	Moralis     Moralis      `toml:"moralis"`
	BscScan     BscScan      `toml:"bscscan"`
	Contract    Contract     `toml:"contract"`
	Redis       Redis        `toml:"redis"`
	Chain       Chain        `toml:"chain"`
	Mysql       Mysql        `toml:"mysql"`
	SignService SignService  `toml:"signService"`
	SignWorkers []SignWorker `toml:"signWorkers"`
}

type Mysql struct {
	Path         string `json:"path" toml:"path"`
	Port         string `json:"port" toml:"port"`
	Config       string `json:"config" toml:"config"`
	Dbname       string `json:"db_name" toml:"dbName"`
	Username     string `json:"username" toml:"username"`
	Password     string `json:"password" toml:"password"`
	MaxIdleConns int    `json:"max_idle_conns" toml:"maxIdleConns"`
	MaxOpenConns int    `json:"max_open_conns" toml:"maxOpenConns"`
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
