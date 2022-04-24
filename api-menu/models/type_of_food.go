package models

import (
	"api-menu/utils"
	"time"
)

type TypeOfFood struct {
	//gorm.Model
	ID   uint32 `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"NOT NULL" json:"name"`

	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Food      []Food     `gorm:"ForeignKey:TypeID"`
}

func CreateTypeOfFood() error {
	var createTypeFood = []TypeOfFood{
		{Name: "Normal"}, {Name: "Vegan"}, {Name: "Vegaterian"},
	}
	utils.MySQLDB.Create(&createTypeFood)
	return nil

}

func CheckType() error {
	type check struct {
		ID uint32
	}
	var typeCheck check
	utils.MySQLDB.Table("type_of_foods").Select("id").Find(&typeCheck)
	if typeCheck.ID == 0 {
		CreateTypeOfFood()
	}
	return nil
}
