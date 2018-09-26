package Model

import (
	"github.com/huacnlee/gobackup/config"
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
		model.perform()
	}
}

func performOne(modelName string) {
	for _, modelConfig := range config.Models {
		if modelConfig.Name == modelName {
			model := Model{
				Config: modelConfig,
			}
			model.perform()
			return
		}
	}
}
