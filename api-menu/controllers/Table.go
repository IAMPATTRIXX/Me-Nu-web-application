package controllers

import (
	"api-menu/models"
	"api-menu/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TableEnpoint(router *gin.RouterGroup){
	router.POST("/create-table", createTable)
	router.GET("/get-table",GetTable)
	router.POST("/update-table", updateTable)
	router.POST("/delte-table", deleteTable)
}
func createTable(c *gin.Context) {
	type RequestData struct {
		TableNO 		uint32 `json:"table_no"`
		Status 			bool 	`json:"status"`
		Url 			string 	`json:"url"`
	}
	var requestTable RequestData
	//var requestData RequestData
	if e := c.BindJSON(&requestTable); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	err := models.CreateTable(
		requestTable.TableNO,
		requestTable.Status,
		requestTable.Url,
	)
	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid ", http.StatusInternalServerError))
		return
	}
	title := "Table"
	description := "create table successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))
	

}

func GetTable(c *gin.Context) {
	Tables, err := models.GetTable()
	if err != nil {
		c.AbortWithStatusJSON(404, utils.ErrorMessage("Menu not fount", http.StatusExpectationFailed))
		return
	}
	type ResponseData struct {
		
		Table interface{} `json:"table"`
	}

	responseData := ResponseData{
		Table: Tables,
	}
	title := "Table"
	description := "Get table successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        responseData,

	}))
}

func updateTable(c *gin.Context) {
	var requestTable *models.Table
	//var requestData RequestData
	if e := c.BindJSON(&requestTable); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	err := models.UpdateTable(requestTable)
	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid ", http.StatusInternalServerError))
		return
	}
	title := "Table"
	description := "update table successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))
}

func deleteTable(c *gin.Context){
	type RequestData struct {
		TableNO uint32 `json:"table_no"`
	}
	var requestData RequestData
	if e := c.BindJSON(&requestData); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	err := models.DeleteTable(
		requestData.TableNO,
	)
	if err != nil{
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid ", http.StatusInternalServerError))
		return
	}
	title := "Table"
	description := "delete table successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))
}