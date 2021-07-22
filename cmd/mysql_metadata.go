package cmd

import (
	"log"
	config "mysql-metadata/configuration"
	"mysql-metadata/internal/cmd"
	"mysql-metadata/internal/repositories"
	"os"
)

func MysqlMetaData() {
	path, err := os.Getwd()
	path = path + "/configuration/properties"
	config.Configuration, err = config.InitConfig("dev", path)
	if err != nil {
		log.Println(err)
	}

	repositories.InitDbConnection()
	cmd.StartApplication()
}
