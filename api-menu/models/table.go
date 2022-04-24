package models

import (
	"api-menu/utils"
	"time"
)

type Table struct {
	ID      uint32 `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	TableNO uint32 `gorm:"NOT NULL" json:"table_no"`
	Status  bool   `gorm:"NOT NULL" json:"status"`
	Url     string `gorm:"NULL" json:"url"`
	TKey 	string  `gorm:"NULL" json:"t_key"`
	//ParLink string `gorm:"NOT NULL" json:"par_link"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func CreateTable(TableNO uint32, Status bool, Url string) error {
	var createTable = Table{
		TableNO: TableNO,
		Status:  Status,
		Url:     Url,
	}
	utils.MySQLDB.Create(&createTable)
	return nil
}
func UpdateTable(Tables *Table) error {
	// var updateTable = Table{
	// 	// ID: ID,
	// 	// TableNO: TableNO,
	// 	Status: Status,
	// 	Url:    Url,
	// }
	//utils.MySQLDB.Model(Table{}).Where("tables.id = ?",ID).Updates(&updateTable)
	if !Tables.Status {
		utils.MySQLDB.Model(Table{}).Where("tables.table_no = ?", Tables.TableNO).Updates(map[string]interface{}{"status": false})

	} else {
		utils.MySQLDB.Model(Table{}).Where("tables.table_no = ?", Tables.TableNO).Updates(Tables)
	}
	return nil
}

func GetTable() (interface{}, error) {
	type SQLResult struct {
		ID      uint32 `json:"id"`
		TableNO uint32 `json:"table_no"`
		Status  bool   `json:"status"`
		Url     string `json:"url"`
		TKey 	string `json:"t_key"`
	}
	var sqlResult []SQLResult
	mysqlQuery := utils.MySQLDB.Table("tables").Select("*")

	if err := mysqlQuery.Find(&sqlResult).Error; err != nil {
		return nil, err
	}
	return sqlResult, nil
}

// func CheckTable(ID uint32, TableNO uint32, Status bool, Url string) error {
// 	type TableCheck struct {
// 		ID uint32 `json:"id"`
// 		// TableNO uint32
// 		// Status bool
// 		// Url string
// 	}
// 	var tcheck TableCheck
// 	utils.MySQLDB.Table("tables").Select("*").Where("tables.id = ?", ID).
// 		Find(&tcheck)

// 	if tcheck.ID == 0 {
// 		CreateTable(TableNO, Status, Url)
// 	} else {
// 		UpdateTable(ID, TableNO, Status, Url)
// 	}
// 	return nil
// }

func DeleteTable(TableNO uint32) error{
	var table Table
	utils.MySQLDB.Where("tables.table_no = ?", TableNO).Delete(&table)
	return nil
}