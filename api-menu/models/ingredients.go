package models

import (
	"api-menu/utils"

	"github.com/gin-gonic/gin"
)

func CreateIngredientsEndpoints(router *gin.RouterGroup) {

}

type Ingredients struct {
	//gorm.Model
	ID     uint32 `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	Name   string `gorm:"NOT NULL" json:"name"`
	FoodID uint32 `gorm:"NOT NULL" json:"food_id"`
	Food   Food   `gorm:"foreignKey:FoodID;OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Source string `gorm:"NULL" json:"source"`
}

func GetFoodID() uint32 {
	type FoodID struct {
		FoodID uint32 `json:"food_id"`
	}
	var foodid FoodID
	sqlQuery := utils.MySQLDB.Table("foods").Select("MAX(foods.id) AS food_id")
	sqlQuery.Find(&foodid)

	return foodid.FoodID
}
func CreateIngredient(Name string, Food_ID uint32, Source string) error {
	var createingredients = Ingredients{
		Name:   Name,
		FoodID: Food_ID,
		Source: Source,
	}
	utils.MySQLDB.Create(&createingredients)
	return nil

}

func UpdateIngredients(Ingredient *Ingredients) error {
	utils.MySQLDB.Model(Ingredients{}).Where("id = ?", Ingredient.ID).Updates(Ingredient)
	return nil
}
