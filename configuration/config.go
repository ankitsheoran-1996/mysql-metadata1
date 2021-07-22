package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DriverName     string
	DataSourceName string
	BucketName     string
}

var Configuration *Config

func InitConfig(env string, path string) (*Config, error) {
	conf := &Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error while loading configurations", err)
		return nil, err
	}
	err = viper.Unmarshal(conf)
	log.Println("Driver Name: ", conf.DriverName)
	log.Println("DataSource Name: ", conf.DataSourceName)
	log.Println("BucketName: ", conf.BucketName)
	return conf, nil
}
