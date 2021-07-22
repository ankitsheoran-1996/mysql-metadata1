package repositories

import (
	"database/sql"
	"fmt"
	config "mysql-metadata/configuration"
)

type Db struct {
	SqlDb *sql.DB
}

var DbConn Db

func InitDbConnection() {
	var err error
	DbConn.SqlDb, err = ConnectDB()
	if err != nil {
		fmt.Println("Db connection failed ", err.Error())
	}
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open(config.Configuration.DriverName, config.Configuration.DataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
