package models

import (
	"fmt"

	"gorm.io/gorm"
)

//DoautoMigration..

func DoautoMigration(dbSrc *gorm.DB) {
	fmt.Println("Start Migration...")
	dbSrc.AutoMigrate(&Food{})
	dbSrc.AutoMigrate(&Ingredients{})
	dbSrc.AutoMigrate(&OrderDetail{})
	dbSrc.AutoMigrate(&TypeOfFood{})
	dbSrc.AutoMigrate(&Billing{})
	dbSrc.AutoMigrate(&Images{})
	dbSrc.AutoMigrate(&Table{})
	dbSrc.AutoMigrate(&Promotion{})
	dbSrc.AutoMigrate(&Review{})

}

//AddForeignKey ...
// func AddForeignKey(dbSrc *gorm.DB) {
// 	// dbSrc.Table("Food").AddForeignKey("type_id","type_of_food(id)","RESTRICT", "RESTRICT")
// 	dbSrc.Model(&Ingredients{}).AddForeignKey("food_id", "food(id)", "RESTRICT", "RESTRICT")
// 	dbSrc.Migrator().CreateConstraint(&Food{}, "TypeID")
// 	//dbSrc.Migrator().CreateConstraint(&TypeOfFood{,})
// }
