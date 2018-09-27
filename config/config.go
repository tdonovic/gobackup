package config

import (
	"fmt"
	"path"
	"time"

	"github.com/spf13/viper"
)

var (
	// Exist Is config file exist
	Exist bool
	// Models configs
	Models []ModelConfig
	// IsTest env
	IsTest bool
	// HomeDir of user
	HomeDir  string
	TempPath string
)

// ModelConfig for special case
type ModelConfig struct {
	Name         string
	DumpPath     string
	CompressWith SubConfig
	EncryptWith  SubConfig
	StoreWith    SubConfig
	Archive      *viper.Viper
	Databases    []SubConfig
	Storages     []SubConfig
	Viper        *viper.Viper
}

// SubConfig sub config info
type SubConfig struct {
	Name  string
	Type  string
	Viper *viper.Viper
}

// loadConfig from:
// - ./gobackup.yml
// - ~/.gobackup/gobackup.yml
// - /etc/gobackup/gobackup.yml
// func init(request string) {
// 	viper.SetConfigType("yaml")
// 	err := viper.ReadInConfig(bytes.NewBuffer([]byte(request)))
// 	if err != nil {
// 		logger.Error("Load gobackup config faild", err)
// 		return
// 	}

// 	Exist = true
// 	Models = []ModelConfig{}
// 	for key := range viper.GetStringMap("models") {
// 		Models = append(Models, loadModel(key))
// 	}

// 	return
// }

func Start() {
	Exist = true
	Models = []ModelConfig{}
	for key := range viper.GetStringMap("models") {
		Models = append(Models, loadModel(key))
	}
}

func loadModel(key string) (model ModelConfig) {
	model.Name = key
	model.DumpPath = path.Join(TempPath, fmt.Sprintf("%d", time.Now().UnixNano()), key)
	model.Viper = viper.Sub("models." + key)

	model.CompressWith = SubConfig{
		Type:  model.Viper.GetString("compress_with.type"),
		Viper: model.Viper.Sub("compress_with"),
	}

	model.EncryptWith = SubConfig{
		Type:  model.Viper.GetString("encrypt_with.type"),
		Viper: model.Viper.Sub("encrypt_with"),
	}

	model.StoreWith = SubConfig{
		Type:  model.Viper.GetString("store_with.type"),
		Viper: model.Viper.Sub("store_with"),
	}

	model.Archive = model.Viper.Sub("archive")

	loadDatabasesConfig(&model)
	loadStoragesConfig(&model)

	return
}

func loadDatabasesConfig(model *ModelConfig) {
	subViper := model.Viper.Sub("databases")
	for key := range model.Viper.GetStringMap("databases") {
		dbViper := subViper.Sub(key)
		model.Databases = append(model.Databases, SubConfig{
			Name:  key,
			Type:  dbViper.GetString("type"),
			Viper: dbViper,
		})
	}
}

func loadStoragesConfig(model *ModelConfig) {
	subViper := model.Viper.Sub("storages")
	for key := range model.Viper.GetStringMap("storages") {
		dbViper := subViper.Sub(key)
		model.Storages = append(model.Storages, SubConfig{
			Name:  key,
			Type:  dbViper.GetString("type"),
			Viper: dbViper,
		})
	}
}

// GetModelByName get model by name
func GetModelByName(name string) (model *ModelConfig) {
	for _, m := range Models {
		if m.Name == name {
			model = &m
			return
		}
	}
	return
}

// GetDatabaseByName get database config by name
func (model *ModelConfig) GetDatabaseByName(name string) (subConfig *SubConfig) {
	for _, m := range model.Databases {
		if m.Name == name {
			subConfig = &m
			return
		}
	}
	return
}
