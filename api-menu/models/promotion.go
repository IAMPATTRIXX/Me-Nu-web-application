package models

import (
	"api-menu/utils"
	"time"
)

type Promotion struct {
	//gorm.Model
	ID        uint32    `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	ImagePro  string    `gorm:"NOT NULL" json:"image_pro"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

func SaveImageProToDB(url string) {
	var saveImage = Promotion{
		ImagePro: url,
	}
	utils.MySQLDB.Create(&saveImage)
}

func GetPromotion() (interface{}, error) {
	type SQLResult struct {
		ImagePro string `json:"image_pro"`
	}
	var sqlResult []SQLResult

	MySQLQuery := utils.MySQLDB.Table("promotions").Select("promotions.image_pro")

	if err := MySQLQuery.Find(&sqlResult).Error; err != nil {
		return nil, err
	}
	return sqlResult, nil
}

func DeletePro(ID uint32) error {
	var promotion Promotion
	utils.MySQLDB.Where("id = ?", ID).Delete(&promotion)
	return nil
}
