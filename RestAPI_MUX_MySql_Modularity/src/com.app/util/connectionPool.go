package util

import (
	"fmt"
	"log"
	"os"

	"com.app/databasestruct"
	"database/sql"
	_ "encoding/json"
	_ "flag"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
	_ "log"
	_ "net/http"
)

var err error

// Config ...
type Config struct {
	DbPassword string
	DbName     string
	DbUsername string
	DbURL      string
}

// Reads info from config file
func ReadConfig() Config {
	var configfile = "D:/apps/content/sandeepDbProperties/sandeepdb.properties"
	fmt.Println(configfile)
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}
	fmt.Println(err)
	var config1 Config
	if _, err := toml.DecodeFile(configfile, &config1); err != nil {
		log.Fatal("Not decoding the config", err)
	}
	//log.Print(config.Index)
	fmt.Println(err)
	fmt.Println(config1)
	return config1
}

func ConnectionStringMain() {
	var config = ReadConfig()
	fmt.Println("In main func", config)
	fmt.Println(config.DbUsername)
	username := config.DbUsername
	password := config.DbPassword
	dbName := config.DbName
	dbURL := config.DbURL
	path := username + ":" + password + "@tcp(" + dbURL + ")/" + dbName + "?parseTime=True"
	fmt.Println(username)
	//  fmt.Printf("%s: %s: %s\n", username, password ,config.DbName)
	fmt.Println("In ConnectionString")
	databasestruct.DBCon, err = sql.Open("mysql", path)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	fmt.Println(databasestruct.DBCon)
	//defer db.Close()

	// make sure connection is available
	db := databasestruct.DBCon
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

}
