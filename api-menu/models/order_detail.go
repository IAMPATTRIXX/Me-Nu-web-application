package models

import (
	"api-menu/utils"
	"fmt"
	"time"
)

type OrderDetail struct {
	//gorm.Model
	ID       uint32 `gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	FoodName string `gorm:"NOT NULL" json:"food_name"`
	Quantity uint32 `gorm:"NOT NULL" json:"quantity"`
	Price    uint32 `gorm:"NOT NULL" json:"price"`
	//TableNO	uint32 	`gorm:"NOT NULL" json:"table_no"`

	//Food	Food	`gorm:"foreignKey:Food_Name;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BillID string `gorm:"NOT NULL" json:"bill_id"`
	IsDone bool   `gorm:"NOT NUL" json:"is_done"`
	//Billing		Billing	`gorm:"foreignKey:BillBID"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func CheckOrder(Order *OrderDetail) error {
	type Orders struct {
		ID        uint32
		BillID    string
		Food_Name string
	}
	var orderCheck Orders
	utils.MySQLDB.Table("order_details").
		Select("*").
		Where("bill_id = ?", Order.BillID).Find(&orderCheck)

	if orderCheck.BillID == "" {
		CreateOrderDetail(Order.FoodName, Order.Price, Order.BillID, Order.Quantity, Order.IsDone)
	} else {
		fmt.Println(Order)
		UpdateOderDetail(Order)
	}
	return nil
}

// admin create ...
func CreateOrderDetail(Food_Name string, Price uint32, BillID string, Quantity uint32, IsDone bool) error {
	var createOrderdetail = OrderDetail{
		FoodName: Food_Name,
		Price:    Price,
		Quantity: Quantity,
		BillID:   BillID,
		IsDone:   IsDone,
	}
	utils.MySQLDB.Create(&createOrderdetail)
	return nil
}

// Get order for Admin
func GetOrderNotDone() (interface{}, error) {
	type SQLResult struct {
		BillID   string   `json:"bill_id"`
		FoodName []string `json:"food_name"`
		TableNO  uint32   `json:"tableNO"`
		Quantity []uint32 `json:"quantity"`
		Done     bool     `json:"done"`
		Is_Done  []bool   `json:"is_done"`
	}
	type FoodQ struct {
		BillID   string `json:"bill_id"`
		FoodName string `json:"food_name"`
		Quantity uint32 `json:"quantity"`
		IsDone   bool   `json:"is_done"`
	}
	var sqlResult1 []FoodQ
	var sqlResult []SQLResult
	var orderResult []SQLResult
	mysqlQuery := utils.MySQLDB.Table("order_details").
		Select("order_details.bill_id, (order_details.food_name) AS food_name, (order_details.quantity) AS quantity, billings.table_no,billings.done,order_details.is_done").
		Joins(`JOIN billings ON order_details.bill_id = billings.id`).
		Where("billings.is_paid = 0 AND billings.done = 0 OR order_details.is_done = 0").
		Group("order_details.bill_id").
		Order("billings.created_at ASC")

	if err := mysqlQuery.Find(&sqlResult).Error; err != nil {
		return nil, err
	}
	mysqlQuery1 := utils.MySQLDB.Table("order_details").
		Select("order_details.food_name,order_details.quantity,order_details.bill_id,order_details.is_done").
		Joins(`JOIN billings ON order_details.bill_id = billings.id`).
		Where("billings.is_paid = 0 AND billings.done = 0 OR order_details.is_done = 0").
		Order("billings.created_at ASC")

	if err := mysqlQuery1.Find(&sqlResult1).Error; err != nil {
		return nil, err
	}

	for _, item := range sqlResult {
		for _, items := range sqlResult1 {
			if items.BillID == item.BillID {
				item.FoodName = append(item.FoodName, items.FoodName)
				item.Quantity = append(item.Quantity, items.Quantity)
				item.Is_Done = append(item.Is_Done, items.IsDone)
			}
		}
		newdata := SQLResult{
			BillID:   item.BillID,
			FoodName: item.FoodName,
			TableNO:  item.TableNO,
			Quantity: item.Quantity,
			Done:     item.Done,
			Is_Done:  item.Is_Done,
		}
		orderResult = append(orderResult, newdata)
	}

	return orderResult, nil
}

func UpdateOderDetail(Order *OrderDetail) error {

	utils.MySQLDB.Model(OrderDetail{}).Where("order_details.bill_id =?", Order.BillID).Updates(Order)
	return nil
}
