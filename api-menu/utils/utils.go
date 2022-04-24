package utils

import (
	"api-menu/utils/database"
	
	"gorm.io/gorm"
)


var MySQLDB *gorm.DB

func InitialDB() error {
	if MySQLDB == nil {
		DB, err := database.NewSQLDatabase()
		if err != nil {
			return err
		}
		MySQLDB = DB
	}
	return nil
}
