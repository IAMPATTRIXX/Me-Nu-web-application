package models

import (
	"api-menu/utils"
	"fmt"
)

type Images struct {
	//gorm.Model
	ID   uint32 `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	ImageName string `gorm:"NOT NULL" json:"image_name"`
	//DirImage string `gorm:"NOT NULL" json:"dir_image"`
	// FoodID   uint32 `gorm:"NOT NULL" json:"food_id"`
	Food []Food `gorm:"foreignKey:ImageID;OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func CreateImage() error {
	var image = Images{}
	utils.MySQLDB.Create(&image)

	return nil
}

func SaveImageToDB(url string) {
	var saveImage = Images{
		ImageName: url,
	}
	utils.MySQLDB.Create(&saveImage)
}
func GetIDImage() uint32 {
	type RequestIID struct {
		ImageID uint32 `json:"image_id"`
	}

	var requestIID RequestIID
	utils.MySQLDB.Table("images").Select("MAX(id)").Find(&requestIID.ImageID)
	fmt.Println(requestIID.ImageID)
	return requestIID.ImageID
}
