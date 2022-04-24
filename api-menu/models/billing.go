package models

import (
	"api-menu/utils"
	"time"
)

type Billing struct {
	//gorm.Model
	ID string `gorm:"primary_key;index" json:"id"`
	//`gorm:"primary_key;index;AUTO_INCREMENT" json:"id"`
	OrderDetail []OrderDetail `gorm:"foreignKey:BillID"`
	// Review      []Review      `gorm:"foreignKey:BillID"`
	// Review []Review `gorm:"foreignKey:BillID"`
	Done bool `gorm:"NOT NULL" json:"done"`

	Amount uint32 `gorm:"NOT NULL" json:"amount"`
	//BID         string        `gorm:"NOT NULL" jspon:"bId"`
	Is_Paid bool   `gorm:"NOT NULL" json:"is_paid"`
	TableNO uint32 `gorm:"NOT NULL" json:"tableNO"`
	//OrderDetail []OrderDetail `gorm:"polymorphic:Bill"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func CheckBill(ID string, Bill *Billing) error {
	type BillCheck struct {
		ID      string
		Amount  uint32
		Is_Paid bool
		TableNO uint32
	}
	var billcheck BillCheck
	utils.MySQLDB.Table("billings").Where("billings.id = ?", ID).
		Find(&billcheck)

	if billcheck.ID == "" {
		CreateBill(ID, Bill)
	} else {
		UpdateBill(Bill, ID)
	}
	return nil
}

func CreateBill(ID string, Bill *Billing) error {

	utils.MySQLDB.Create(&Bill)
	return nil

}

func UpdateBill(Bill *Billing, ID string) error {

	utils.MySQLDB.Model(Billing{}).Where("billings.id = ?", ID).Updates(Bill)
	return nil

}

func GetBillNotPaid() (interface{}, error) {

	type SQLOrders struct {
		BillID   string `json:"bill_id"`
		FoodName string `json:"food_name"`
		Quantity uint32 `json:"quantity"`
		Price    uint32 `json:"price"`
	}

	type SQLResult struct {
		ID        string      `json:"id"`
		TableNO   uint32      `json:"table_no"`
		CreatedAt time.Time   `json:"created_at"`
		Amount    uint32      `json:"amount"`
		IsPaid    bool        `json:"is_paid"`
		OrderBill []SQLOrders `gorm:"-" json:"order_bill"`
	}
	var orderBill []SQLOrders
	var billResult []SQLResult
	var sqlResult []SQLResult

	mysqlQuery := utils.MySQLDB.Table("billings").
		Select("billings.id,billings.table_no ,billings.created_at ,billings.amount,billings.is_paid").
		//Joins(`JOIN billings ON oreder_details.b_id = billings.b_id`).
		Where(" billings.done = 1 AND billings.is_paid = 0").
		Order("billings.updated_at ASC")
	if err := mysqlQuery.Find(&sqlResult).Error; err != nil {
		return nil, err
	}

	mysqlQuery1 := utils.MySQLDB.Table("order_details").Select("order_details.food_name, order_details.quantity, order_details.price,order_details.bill_id").
		Where("order_details.is_done = 1")
	if err := mysqlQuery1.Find(&orderBill).Error; err != nil {
		return nil, err
	}

	for _, item := range sqlResult {
		for _, items := range orderBill {
			if item.ID == items.BillID {
				item.OrderBill = append(item.OrderBill, items)
			}
		}
		newdata := SQLResult{
			ID:        item.ID,
			TableNO:   item.TableNO,
			CreatedAt: item.CreatedAt,
			Amount:    item.Amount,
			IsPaid:    item.IsPaid,
			OrderBill: item.OrderBill,
		}
		billResult = append(billResult, newdata)
	}
	return billResult, nil
}
func GetBillIsPaid() (interface{}, error) {
	type SQLOrders struct {
		FoodName string `json:"food_name"`
		Quantity uint32 `json:"quantity"`
		Price    uint32 `json:"price"`
		BillID   string `json:"bill_id"`
	}
	type SQLResult struct {
		ID        string      `json:"id"`
		TableNO   uint32      `json:"table_no"`
		UpdatedAt time.Time   `json:"updated_at"`
		Amount    uint32      `json:"amount"`
		Is_Paid   bool        `json:"is_paid"`
		OrderBill []SQLOrders `gorm:"-" json:"order_bill"`
	}
	var billResult []SQLResult
	var sqlResult []SQLResult
	mysqlQuery := utils.MySQLDB.Table("billings").
		Select("billings.id,billings.table_no ,billings.updated_at ,billings.amount,billings.is_paid").
		//Joins(`JOIN billings ON oreder_details.b_id = billings.b_id`).
		Where(" billings.done = 1 AND billings.is_paid = 1").
		Order("billings.updated_at ASC")
	if err := mysqlQuery.Find(&sqlResult).Error; err != nil {
		return nil, err
	}

	var orderbill []SQLOrders
	mysqlQuery1 := utils.MySQLDB.Table("order_details").Select("order_details.food_name, order_details.quantity, order_details.price,order_details.bill_id").
		Where("order_details.is_done = 1")
	if err := mysqlQuery1.Find(&orderbill).Error; err != nil {
		return nil, err
	}
	for _, item := range sqlResult {
		for _, items := range orderbill {
			if item.ID == items.BillID {
				item.OrderBill = append(item.OrderBill, items)
			}
		}
		newdata := SQLResult{
			ID:        item.ID,
			TableNO:   item.TableNO,
			UpdatedAt: item.UpdatedAt,
			Amount:    item.Amount,
			Is_Paid:   item.Is_Paid,
			OrderBill: item.OrderBill,
		}
		billResult = append(billResult, newdata)
	}
	return billResult, nil
}
