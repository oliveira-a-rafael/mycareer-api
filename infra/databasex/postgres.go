package databasex

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/oliveira-a-rafael/mycareer-api/config"
)

type Repository struct {
	dbPostgres *gorm.DB
}

func GetInstancePostgres() *gorm.DB {

	var (
		env            = config.APP_ENV
		username       = config.DB_USER
		password       = config.DB_PASS
		dbName         = config.DB_NAME
		port           = config.DB_PORT
		host           = config.DB_HOST
		connectionName = config.SQL_CONN_NAME
		conec          string
	)

	if env == "ALTER" {
		conec = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, username, dbName, password)
	} else {
		if env == "PROD" {
			conec = fmt.Sprintf("user=%s password=%s host=/cloudsql/%s dbname=%s", username, password, connectionName, dbName)
		} else {
			conec = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbName, password)
		}
	}

	fmt.Println("connect: ", conec)

	dbPostgres, err := gorm.Open("postgres", conec)

	if err != nil {
		panic(fmt.Sprintf("Error to connect DB: username=%s db_name=%s port=%s host=%s detail_erro=%s", username, dbName, port, host, err.Error()))
	}

	dbPostgres.LogMode(true)

	return dbPostgres
}
