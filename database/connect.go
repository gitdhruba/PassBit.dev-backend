package database

//This package handles the connection of external postgres DB
//Author : Dhruba Sinha

import (
	"fmt"
	"os"
	"passbit/config"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB //db instance for exporting

// following function connects to the db and assigns DB
func ConnectDB() {

	var err error
	var port int

	port, err = strconv.Atoi(config.Config("DBPORT"))
	if err != nil {
		fmt.Println("ERROR : invalid DBPORT")
		os.Exit(1)
	}

	//data-source-name for connecting to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", config.Config("DBHOST"), port, config.Config("DBUNAME"), config.Config("DBPASSWORD"), config.Config("DBNAME"))

	//open db connection
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("ERROR : could not connect to database")
		os.Exit(1)
	}

	fmt.Println("Connection Opened to Database")

	//set DB logger
	DB.Logger = logger.Default.LogMode(logger.Info)

}
