package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"spike-frame/constant"
)

func InitViper() *viper.Viper {
	var configPath string
	if configEnv := os.Getenv(constant.ConfigEnv); configEnv == "" {
		configPath = constant.ConfigFile
		log.Infof("use default config path: %v", configPath)
	} else {
		configPath = configEnv
		log.Infof("use env config path %v\n", configPath)
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("toml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Info("config file changed:", e.Name)
		if err := v.Unmarshal(&Cfg); err != nil {
			log.Error(err)
		}
		log.Infof("cfg: %+v", Cfg)
	})
	if err := v.Unmarshal(&Cfg); err != nil {
		log.Error(err)
	}
	log.Infof("cfg: %+v", Cfg)
	return v
}
