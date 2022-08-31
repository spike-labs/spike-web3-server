package util

import (
	logger "github.com/ipfs/go-log"
	"os"
)

var log = logger.Logger("util")

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			log.Infof("create directory : %s", v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				log.Infof("create directory : %s", err.Error())
				return err
			}
		}
	}
	return err
}
