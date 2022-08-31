package config

import (
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("config")

var Cfg Config

type Config struct {
	System   System   `toml:"system"`
	Moralis  Moralis  `toml:"moralis"`
	BscScan  BscScan  `toml:"bscscan"`
	Contract Contract `toml:"contract"`
	Redis    Redis    `toml:"redis"`
	Chain    Chain    `toml:"chain"`
	Mysql    Mysql    `toml:"mysql"`
}

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" toml:"path"`
	Port         string `mapstructure:"port" json:"port" toml:"port"`
	Config       string `mapstructure:"config" json:"config" toml:"config"`
	Dbname       string `mapstructure:"db_name" json:"db_name" toml:"dbame"`
	Username     string `mapstructure:"username" json:"username" toml:"username"`
	Password     string `mapstructure:"password" json:"password" toml:"password"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" toml:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" toml:"maxOpenConns"`
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
