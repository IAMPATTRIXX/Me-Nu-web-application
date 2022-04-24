package models

import (
	"api-menu/utils"
	"time"
)

type Review struct {
	//gorm.Model
	ID uint32 `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`

	Reviews  string `gorm:"NOT NULL" json:"reviews"`
	FoodName string `gorm:"NOT NULL" json:"food_name"`
	BillID   string `gorm:"NOT NULL" json:"bill_id"`

	//BillID string  `gorm:"NOT NULL" json:"bill_id"`

	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

// Customer ...
func CreateReview(Reviews string, Food_Name string, BillID string) error {

	var createReview = Review{
		Reviews:  Reviews,
		FoodName: Food_Name,
		BillID:   BillID,
	}
	utils.MySQLDB.Create(&createReview)
	return nil

}

// Admin ...
func deleteReview(ID uint32) error {
	utils.MySQLDB.Delete(&Review{}, ID)
	return nil
}
