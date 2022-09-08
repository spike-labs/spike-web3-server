package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
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
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Info("config file changed:", e.Name)

		file, err := os.Open(configPath)
		switch {
		case os.IsNotExist(err):
			log.Errorf("%v", err)
			return
		case err != nil:
			return
		}
		_, err = toml.NewDecoder(file).Decode(&Cfg)
		if err != nil {
			return
		}
		log.Infof("cfg: %+v", Cfg)
	})

	file, err := os.Open(configPath)
	switch {
	case os.IsNotExist(err):
		panic("config is not exist")
	case err != nil:
		panic("config path error")
	}
	_, err = toml.NewDecoder(file).Decode(&Cfg)
	if err != nil {
		panic("init config err")
	}

	log.Infof("cfg: %+v", Cfg)
	return v
}
