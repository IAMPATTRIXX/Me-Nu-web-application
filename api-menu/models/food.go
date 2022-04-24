package models

import (
	"api-menu/utils"
	"time"
)

type Food struct {
	//gorm.Model
	ID uint32 `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	//Images      []Images    `gorm:"foreignKey:FoodID"`
	//Review []Review `gorm:"foreignKey:FoodID"`
	Name string `gorm:"NOT NULL" json:"name"`
	//Ingredients Ingredients `gorm:"foreignKey:FoodID"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`

	SpcFood   bool       `gorm:"NULL" json:"spc_food"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Price     uint32     `gorm:"NOT NULL" json:"price"`
	TypeID    uint32     `gorm:"NOT NULL" json:"type_id" `
	//TypeOfFood	TypeOfFood `gorm:"foreignKey:TypeID"`
	//TypeOfFood TypeOfFood `gorm:"foreignKey:TypeID "`

	Has_Milk    bool   `gorm:"NOT nuLL" json:"has_milk"`
	Has_Egg     bool   `gorm:"NOT nuLL" json:"has_egg"`
	Has_Bean    bool   `gorm:"NOT nuLL" json:"has_bean"`
	Has_Seafood bool   `gorm:"NOT nuLL" json:"has_seafood"`
	Status      bool   `gorm:"NOT NULL" json:"status"`
	Description string `gorm:"NULL" json:"description"`
	ImageID     uint32 `gorm:"NULL" json:"image_id"`
}

//Customer
func GetFoodCus() (interface{}, error) {
	type SQLIngerdients struct {
		ID     uint32 `json:"id"`
		FoodID uint32 `json:"food_id"`
		Name   string `json:"name"`
		Source string `json:"source"`
	}
	type SQLResult struct {
		ID          uint32 `json:"id"`
		Name        string `json:"name"`
		TypeID      uint32 `json:"type_id"`
		Status      bool   `json:"status"`
		Price       uint32 `json:"price"`
		SpcFood     bool   `json:"spc_food"`
		Has_Milk    bool   `json:"has_milk"`
		Has_Egg     bool   `json:"has_egg"`
		Has_Bean    bool   `json:"has_bean"`
		Has_Seafood bool   `json:"has_seafood"`
		Description string `json:"description"`
		ImageName   string `json:"image_name"`
		// Reviews     []string         `json:"reviews"`
		// RCreatedAt  []string         `json:"rcreated_at"`
		Review  []Review         `gorm:"-" json:"review"`
		IngName []SQLIngerdients `gorm:"-" json:"ing_name"`
	}
	type SQLResults struct {
		FoodName   string `json:"food_name"`
		Reviews    string `json:"reviews"`
		RCreatedAt string `json:"rcreated_at"`
	}

	//var sqlResults []SQLResults
	var resultFood []SQLResult
	var sqlResults []Review
	var resultIng []SQLIngerdients
	mysqlQuery := utils.MySQLDB.Table("foods").Select("*").Joins("JOIN images ON foods.image_id = images.id ").Where("foods.status = 1")

	if err := mysqlQuery.Find(&resultFood).Error; err != nil {
		return nil, err
	}
	mysqlQuery2 := utils.MySQLDB.Table("reviews").
		Select("*")
	if err := mysqlQuery2.Find(&sqlResults).Error; err != nil {
		return nil, err
	}

	mysqlQuery3 := utils.MySQLDB.Table("ingredients").
		Select("*")
	if err := mysqlQuery3.Find(&resultIng).Error; err != nil {
		return nil, err
	}

	// fmt.Println(sqlResults)
	// var re []qstring
	var foodResults []SQLResult
	for _, item := range resultFood {

		for _, items := range sqlResults {

			if item.Name == items.FoodName {
				item.Review = append(item.Review, items)
				// re = append(re, "pp")
				//item.RCreatedAt = append(item.RCreatedAt, items.RCreatedAt)
				// review := items.Reviews
			}
		}
		for _, itemIn := range resultIng {
			if item.ID == itemIn.FoodID {
				item.IngName = append(item.IngName, itemIn)
				// re = append(re, "pp")
				//item.Source = append(item.Source, itemIn.Source)
				// review := items.Reviews
			}
		}
		newdata := SQLResult{
			ID:          item.ID,
			Name:        item.Name,
			TypeID:      item.TypeID,
			Status:      item.Status,
			Price:       item.Price,
			SpcFood:     item.SpcFood,
			Has_Milk:    item.Has_Milk,
			Has_Egg:     item.Has_Egg,
			Has_Bean:    item.Has_Bean,
			Has_Seafood: item.Has_Seafood,
			Description: item.Description,
			ImageName:   item.ImageName,
			// Reviews:     item.Reviews,
			// RCreatedAt:  item.RCreatedAt,
			Review:  item.Review,
			IngName: item.IngName,
			//Source: item.Source,
		}
		foodResults = append(foodResults, newdata)

	}
	// sqlResultMapper := make(map[uint32]*getFood)
	// for index := range resultFood {
	// 	keyID := resultFood[index].ID
	// 	sqlResultMapper[keyID] = &resultFood[index]
	// }

	return foodResults, nil
}

// Resturant
func CreateFood(Name string, Price uint32, Type_ID uint32, Has_Milk bool, Has_Egg bool, Has_Bean bool, Has_Seafood bool, SpcFood bool, Status bool, description string, ImageID uint32) error {

	var createFood = Food{
		Name:        Name,
		Price:       Price,
		TypeID:      Type_ID,
		Has_Milk:    Has_Milk,
		Has_Egg:     Has_Egg,
		Has_Bean:    Has_Bean,
		Has_Seafood: Has_Seafood,
		SpcFood:     SpcFood,
		Status:      Status,

		Description: description,
		ImageID:     ImageID,
	}
	utils.MySQLDB.Create(&createFood)
	return nil

}

func UpdateFood(ID uint32, Name string, Price uint32, TypeID uint32, Milk bool, Egg bool, Bean bool, Seafood bool, SpcFood bool, Status bool, Description string, ImageID uint32) error {
	var updateFood = Food{
		ID:          ID,
		Name:        Name,
		Price:       Price,
		TypeID:      TypeID,
		Has_Milk:    Milk,
		Has_Egg:     Egg,
		Has_Bean:    Bean,
		Has_Seafood: Seafood,
		Description: Description,
		ImageID:     ImageID,
		SpcFood:     SpcFood,
	}
	utils.MySQLDB.Model(Food{}).Where("foods.id = ?", ID).Updates(&updateFood)
	return nil
}
func ToggleMenu(ID uint32, Status bool) error {
	// var updateStatus = Food{
	// 	Status:   Status,
	// 	Has_Milk: Has_milk,
	// }
	// fmt.Println(updateStatus)
	utils.MySQLDB.Model(Food{}).Where("foods.id = ?", ID).Update("status", Status)
	return nil
}

// DeleteFood ...
func DeleteFood(FoodID uint32) error {
	var err error
	//trx := utils.MySQLDB.Begin()
	type Image struct {
		ID        uint32
		ImageName string
	}
	var foods Food
	var images Image
	err = utils.MySQLDB.Begin().
		Joins("JOIN images ON images.id = foods.image_id ").
		Where("foods.id = ?", FoodID).Find(&foods).Error
	if err != nil {
		utils.MySQLDB.Begin().Rollback()
		return err
	}
	err = utils.MySQLDB.Begin().
		Joins("JOIN foods ON images.id = foods.image_id ").
		Where("foods.id = ?", FoodID).Find(&images).Error
	if err != nil {
		utils.MySQLDB.Begin().Rollback()
		return err
	}
	var ingredients []Ingredients
	err = utils.MySQLDB.Begin().Where("ingredients.food_id = ?", FoodID).Find(&ingredients).Error
	if err != nil {
		utils.MySQLDB.Begin().Rollback()
		return err
	}

	err = utils.MySQLDB.Begin().Exec(`SET FOREIGN_KEY_CHECKS=0`).Error
	if err != nil {
		utils.MySQLDB.Begin().Rollback()
		return err
	}
	for index := range ingredients {
		if ingredients[index].FoodID == foods.ID {
			utils.MySQLDB.Where(`ingredients.food_id = ?`, FoodID).Delete(&ingredients)
		}

		//trx.Exec(`DELETE FROM ingredients WHERE ingredients.food_id = ?`, FoodID)

	}

	//trx.Debug().Exec(`DELETE FROM images WHERE images.id = ?`,foods.ID)
	utils.MySQLDB.Where("foods.id = ?", FoodID).Delete(&foods)
	//trx.Debug().Exec(`DELETE FROM foods WHERE foods.id = ?`,FoodID)
	utils.MySQLDB.Where("images.id = ?", foods.ImageID).Delete(&images)
	return err
}

func GetRakingMenu() (interface{}, error) {

	type SQLIngerdients struct {
		ID     uint32 `json:"id"`
		FoodID uint32 `json:"food_id"`
		Name   string `json:"name"`
		Source string `json:"source"`
	}
	type SQLResult struct {
		ID          uint32 `json:"id"`
		Name        string `json:"name"`
		TypeID      uint32 `json:"type_id"`
		Status      bool   `json:"status"`
		Price       uint32 `json:"price"`
		SpcFood     bool   `json:"spc_food"`
		Has_Milk    bool   `json:"has_milk"`
		Has_Egg     bool   `json:"has_egg"`
		Has_Bean    bool   `json:"has_bean"`
		Has_Seafood bool   `json:"has_seafood"`
		Description string `json:"description"`
		ImageName   string `json:"image_name"`
		// Reviews     []string         `json:"reviews"`
		// RCreatedAt  []string         `json:"rcreated_at"`
		Review  []Review         `gorm:"-" json:"review"`
		IngName []SQLIngerdients `gorm:"-" json:"ing_name"`
	}
	type SQLResults struct {
		FoodName   string `json:"food_name"`
		Reviews    string `json:"reviews"`
		RCreatedAt string `json:"rcreated_at"`
	}
	//var sqlResults []SQLResults
	var sqlResults []Review
	var resultFood []SQLResult
	var resultIng []SQLIngerdients
	mysqlQuery := utils.MySQLDB.Table("foods").Select("*").Joins("JOIN images ON foods.image_id = images.id ").Where("foods.status = 1")

	if err := mysqlQuery.Find(&resultFood).Error; err != nil {
		return nil, err
	}
	mysqlQuery2 := utils.MySQLDB.Table("reviews").
		Select("*")
	if err := mysqlQuery2.Find(&sqlResults).Error; err != nil {
		return nil, err
	}

	mysqlQuery3 := utils.MySQLDB.Table("ingredients").
		Select("*")
	if err := mysqlQuery3.Find(&resultIng).Error; err != nil {
		return nil, err
	}

	// fmt.Println(sqlResults)
	// var re []qstring
	var foodResults []SQLResult
	for _, item := range resultFood {

		for _, items := range sqlResults {

			if item.Name == items.FoodName {
				item.Review = append(item.Review, items)
				// re = append(re, "pp")
				//item.RCreatedAt = append(item.RCreatedAt, items.RCreatedAt)
				// review := items.Reviews
			}
		}
		for _, itemIn := range resultIng {
			if item.ID == itemIn.FoodID {
				item.IngName = append(item.IngName, itemIn)
				// re = append(re, "pp")
				//item.Source = append(item.Source, itemIn.Source)
				// review := items.Reviews
			}
		}
		newdata := SQLResult{
			ID:          item.ID,
			Name:        item.Name,
			TypeID:      item.TypeID,
			Status:      item.Status,
			Price:       item.Price,
			SpcFood:     item.SpcFood,
			Has_Milk:    item.Has_Milk,
			Has_Egg:     item.Has_Egg,
			Has_Bean:    item.Has_Bean,
			Has_Seafood: item.Has_Seafood,
			Description: item.Description,
			ImageName:   item.ImageName,
			// Reviews:     item.Reviews,
			// RCreatedAt:  item.RCreatedAt,
			Review:  item.Review,
			IngName: item.IngName,
			//Source: item.Source,
		}
		foodResults = append(foodResults, newdata)

	}

	type SQLRank struct {
		Name   string      `json:"name"`
		Amount uint32      `json:"amount"`
		Food   []SQLResult `gorm:"-" json:"food"`
	}

	t := time.Now()
	lastT := t.AddDate(0, 0, -14)
	// currentT := t.AddDate(0,0,+1)
	// currentTT := currentT.Format("2006-01-02")
	currentT := t.Format("2006-01-02")

	fromDate := lastT.Format("2006-01-02")
	toDate := currentT

	var sqlResult []SQLRank
	var sqlRanking []SQLRank
	mysqlQuery4 := utils.MySQLDB.Table("order_details").
		Select(`order_details.food_name AS name , SUM(order_details.quantity) AS amount`).
		Joins("LEFT JOIN foods ON foods.name = order_details.food_name ").
		Where("foods.status = 1 AND order_details.created_at >= ?  AND order_details.created_at < ? ", fromDate, toDate).
		Group("order_details.food_name").Order("amount DESC")
	//Where("foods.name = ord.food_name")
	//Group("ord.food_name")
	if err := mysqlQuery4.Find(&sqlResult).Error; err != nil {
		return nil, err
	}

	for _, item := range sqlResult {
		for _, items := range foodResults {
			if item.Name == items.Name {
				item.Food = append(item.Food, items)
			}
		}
		newdata := SQLRank{
			Name:   item.Name,
			Amount: item.Amount,
			Food:   item.Food,
		}
		sqlRanking = append(sqlRanking, newdata)
	}

	return sqlRanking, nil
}

func GetFoodRes() (interface{}, error) {
	type SQLIngerdients struct {
		ID     uint32 `json:"id"`
		FoodID uint32 `json:"food_id"`
		Name   string `json:"name"`
		Source string `json:"source"`
	}
	type SQLResult struct {
		ID          uint32 `json:"id"`
		Name        string `json:"name"`
		TypeID      uint32 `json:"type_id"`
		Status      bool   `json:"status"`
		Price       uint32 `json:"price"`
		SpcFood     bool   `json:"spc_food"`
		Has_Milk    bool   `json:"has_milk"`
		Has_Egg     bool   `json:"has_egg"`
		Has_Bean    bool   `json:"has_bean"`
		Has_Seafood bool   `json:"has_seafood"`
		Description string `json:"description"`
		ImageName   string `json:"image_name"`
		//Reviews     []string         `json:"reviews"`
		//RCreatedAt  []string         `json:"rcreated_at"`
		Review  []Review         `gorm:"-" json:"review"`
		IngName []SQLIngerdients `gorm:"-" json:"ing_name"`
	}
	type SQLResults struct {
		FoodName  string `json:"food_name"`
		Reviews   string `json:"reviews"`
		CreatedAt string `json:"createdAt"`
	}

	var resultFood []SQLResult
	//var sqlResults []SQLResults
	var sqlResults []Review
	var resultIng []SQLIngerdients
	mysqlQuery := utils.MySQLDB.Table("foods").Select("*").
		Joins("JOIN images ON foods.image_id = images.id")
		//Joins("LEFT JOIN (SELECT JSON_ARRAYAGG(reviews) AS reviews,JSON_ARRAYAGG(created_at) AS rcreated_at,reviews.food_id FROM reviews GROUP BY reviews.food_id) AS review ON foods.id = review.food_id ")
		//Group("foods.id")

	if err := mysqlQuery.Find(&resultFood).Error; err != nil {
		return nil, err
	}

	mysqlQuery2 := utils.MySQLDB.Table("reviews").
		Select("*")
	if err := mysqlQuery2.Find(&sqlResults).Error; err != nil {
		return nil, err
	}
	mysqlQuery3 := utils.MySQLDB.Table("ingredients").
		Select("*")
	if err := mysqlQuery3.Find(&resultIng).Error; err != nil {
		return nil, err
	}
	// fmt.Println(sqlResults)
	// var re []qstring
	var foodResults []SQLResult
	for _, item := range resultFood {

		for _, items := range sqlResults {

			if item.Name == items.FoodName {
				item.Review = append(item.Review, items)
				// re = append(re, "pp")
				//item.RCreatedAt = append(item.RCreatedAt, items.RCreatedAt)
				// review := items.Reviews
			}
		}
		for _, itemIn := range resultIng {
			if item.ID == itemIn.FoodID {
				item.IngName = append(item.IngName, itemIn)
				// re = append(re, "pp")

				// review := items.Reviews
			}
		}
		newdata := SQLResult{
			ID:          item.ID,
			Name:        item.Name,
			TypeID:      item.TypeID,
			Status:      item.Status,
			Price:       item.Price,
			SpcFood:     item.SpcFood,
			Has_Milk:    item.Has_Milk,
			Has_Egg:     item.Has_Egg,
			Has_Bean:    item.Has_Bean,
			Has_Seafood: item.Has_Seafood,
			Description: item.Description,
			ImageName:   item.ImageName,
			// Reviews:     item.Reviews,
			// RCreatedAt:  item.RCreatedAt,
			Review:  item.Review,
			IngName: item.IngName,
		}
		foodResults = append(foodResults, newdata)

	}

	return foodResults, nil
}
