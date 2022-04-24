package controllers

import (
	"api-menu/models"
	"api-menu/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateReviewEndpoints(router *gin.RouterGroup) {
	router.POST("/create-review",createReview)
}
func createReview(c *gin.Context) {
	type RequestData struct {
		Reviews string `json:"reviews"`
		FoodName string `json:"food_name"`
		BillID string `json:"bill_id"`
	}

	var requestData RequestData

	if e := c.BindJSON(&requestData); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	err := models.CreateReview(
		requestData.Reviews,
		requestData.FoodName,
		requestData.BillID,
	)
	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid ", http.StatusInternalServerError))
		return
	}
	title := "Review"
	description := "create review successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))
	return
}
