package gobackup

import (
	"bytes"

	"github.com/huacnlee/gobackup/logger"
	"github.com/spf13/viper"
	"github.com/tdonovic/gobackup/config"
)

const (
	usage = "Easy full stack backup operations on UNIX-like systems"
)

var (
	modelName = ""
	version   = "master"
)

func performAll() {
	for _, modelConfig := range config.Models {
		model := Model{
			Config: modelConfig,
		}
		perform()
	}
}

func performOne(request string) {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(request)))
	if err != nil {
		logger.Error("Load gobackup config faild", err)
		return
	} else {
		config.Start()
	}

	for _, modelConfig := range config.Models {
		if modelConfig.Name == modelName {
			model := Model{
				Config: modelConfig,
			}
			perform()
			return
		}
	}
}
