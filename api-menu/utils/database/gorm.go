package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"

	"github.com/tidwall/gjson"
	_ "github.com/tidwall/gjson"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseConfiguration struct {
	Username string
	Password string
	Hostname string
	Port string
	Database string
}


func NewSQLDatabase() (*gorm.DB, error) {

	var err error
	var target string = "Dev1"
	body, _ := ioutil.ReadFile("./configs/json/mysql_database.json")
	configValue := gjson.Get(string(body), target)
	currentConfig := databaseConfiguration{}
	json.Unmarshal([]byte(configValue.String()), &currentConfig)

	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True",
	currentConfig.Username,
	currentConfig.Password,
	currentConfig.Hostname,
	currentConfig.Port,
	currentConfig.Database,
	)
	MySQL, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)

	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return MySQL, nil
}

