package configs

import (
	"log"

	"github.com/spf13/viper"
)

type store struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type config struct {
	Environment string `json:"environment"`
	Port        string `json:"port"`
	Redis       store  `json:"store"`
	Memcache    store  `json:"memcache"`
}

var Configuration config

func InitConfigurations() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("err: ", err.Error())
		return err
	}

	var cnf config
	err = viper.Unmarshal(&cnf)
	if err != nil {
		log.Println("err: ", err.Error())
		return err
	}

	Configuration = cnf
	return nil
}
