package controllers

import (
	"api-menu/models"
	"api-menu/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBillingEndpoints(router *gin.RouterGroup) {
	router.POST("/create-bill", createBill)
	router.POST("/create-order", createOrder)
	router.GET("/get-order", getOrderNotDone)
	router.GET("/get-bill", getBill)
	router.POST("/update-order", updateOrder)
}

func createBill(c *gin.Context) {
	type RequestData struct {
		ID      string `json:"id"`
		Amount  uint32 `json:"amount"`
		Is_Paid bool   `json:"is_paid"`
		TableNO uint32 `json:"table_no"`
		Done    bool   `json:"done"`
	}
	var requestD *models.Billing

	if e := c.BindJSON(&requestD); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input1", http.StatusBadRequest))
		return
	}
	err := models.CheckBill(

		requestD.ID,
		requestD)

	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid ", http.StatusInternalServerError))
		return
	}
	title := "Bill"
	description := "create bill successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))

}

func createOrder(c *gin.Context) {
	type RequestData struct {
		Food_Name string `json:"food_name"`
		Price     uint32 `json:"price"`
		BillID    string `json:"bill_id"`
		Quantity  uint32 `json:"quantity"`
		IsDone    bool   `json:"is_done"`
	}
	var requestData []RequestData

	if e := c.BindJSON(&requestData); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	for index := range requestData {
		models.CreateOrderDetail(
			requestData[index].Food_Name,
			requestData[index].Price,
			requestData[index].BillID,
			requestData[index].Quantity,
			requestData[index].IsDone,
		)
	}

	title := "Order"
	description := "create order successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))

}

func updateOrder(c *gin.Context) {
	var requestD models.OrderDetail
	if e := c.BindJSON(&requestD); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	models.UpdateOderDetail(&requestD)
	title := "Order"
	description := "update order successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))
}

func getOrderNotDone(c *gin.Context) {
	results, err := models.GetOrderNotDone()
	if err != nil {
		c.AbortWithStatusJSON(404, utils.ErrorMessage("Order not fount", http.StatusExpectationFailed))
		return
	}
	type ResponseData struct {
		Order interface{} `json:"order"`
	}
	responseData := ResponseData{
		Order: results,
	}
	title := "Order"
	description := "Get order successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        responseData,
	}))
}

func getBill(c *gin.Context) {
	billnotpaid, err := models.GetBillNotPaid()
	if err != nil {
		c.AbortWithStatusJSON(404, utils.ErrorMessage("Bill not fount", http.StatusExpectationFailed))
		return
	}
	billpaid, err := models.GetBillIsPaid()
	if err != nil {
		c.AbortWithStatusJSON(404, utils.ErrorMessage("Bill not fount", http.StatusExpectationFailed))
		return
	}
	type ResponseData struct {
		BillPaid  interface{} `json:"bill_paid"`
		BillNPaid interface{} `json:"bill_npaid"`
	}
	responseData := ResponseData{
		BillNPaid: billnotpaid,
		BillPaid:  billpaid,
	}
	title := "Bill"
	description := "Get bill paided successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        responseData,
	}))
}
