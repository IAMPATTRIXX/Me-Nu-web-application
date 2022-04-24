package controllers

import (
	"api-menu/models"
	"api-menu/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFoodEndpoints(router *gin.RouterGroup) {
	router.GET("/res/menu-food", getMenuRes)
	router.POST("/create-food", createFood)
	router.POST("/delete-food", deleteFood)
	router.GET("/cus/menu-food", getMenuCus)
	router.POST("/update-food", updateFood)
	router.POST("/update-ingredients", updateIngredients)
	router.POST("/create-ingredients", createingredients)
	router.POST("/status/toggle", toggleMenu)
	router.POST("/delete-pro",deletePro)
	//http.HandleFunc("/upload", up)
	//router.POST("/upload-image", FirebaseStorage.UploadImage1)
	//router.POST("/up",TCre)
	//http.HandleFunc("/up/image", UploadImage)

}

func getMenuRes(c *gin.Context) {
	results, err := models.GetFoodRes()

	if err != nil {
		c.AbortWithStatusJSON(404, utils.ErrorMessage("Menu not fount", http.StatusExpectationFailed))
		return
	}
	type ResponseData struct {
		Food interface{} `json:"food"`
	}

	responseData := ResponseData{
		Food: results,
	}
	title := "Menu Food"
	description := "Get menu successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        responseData,
		//models.GetImage(c.Writer),

	}))

	//models.GetImage(c.Writer)
	//fmt.Fprintln(c.Writer)
	// create type ...
	models.CheckType()

}

func getMenuCus(c *gin.Context) {
	results, err := models.GetFoodCus()
	if err != nil {
		c.AbortWithStatusJSON(404, utils.ErrorMessage("Menu not fount", http.StatusExpectationFailed))
		return
	}
	ranking, err := models.GetRakingMenu()
	if err != nil {
		c.AbortWithStatusJSON(404, utils.ErrorMessage("Ranking menu not fount", http.StatusExpectationFailed))
		return
	}
	promotion, err := models.GetPromotion()
	if err != nil {
		c.AbortWithStatusJSON(404, utils.ErrorMessage("Promotion not fount", http.StatusExpectationFailed))
		return
	}
	type ResponseData struct {
		Food      interface{} `json:"food"`
		Ranking   interface{} `json:"ranking"`
		Promotion interface{} `json:"promotion"`
	}

	responseData := ResponseData{
		Food:      results,
		Ranking:   ranking,
		Promotion: promotion,
	}
	title := "Menu Food"
	description := "Get menu successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        responseData,
		//models.GetImage(c.Writer),

	}))

	models.CheckType()

}

func createFood(c *gin.Context) {
	type RequestData struct {
		ID          uint32 `json:"id"`
		Name        string `json:"name"`
		Price       uint32 `json:"price"`
		TypeID      uint32 `json:"type_id"`
		Has_Milk    bool   `json:"has_milk"`
		Has_Egg     bool   `json:"has_egg"`
		Has_Bean    bool   `json:"has_bean"`
		Has_Seafood bool   `json:"has_seafood"`
		Status      bool   `json:"status"`
		Description string `json:"description"`
		ImageID     uint32 `json:"image_id"`
		SpcFood     bool   `json:"spc_food"`
		//image *http.Request `json:"image"`
		NewImg bool `json:"new_img"`
	}

	var requestData RequestData

	if e := c.BindJSON(&requestData); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	var ImageID = models.GetIDImage()
	err := models.CreateFood(

		requestData.Name,
		requestData.Price,
		requestData.TypeID,
		requestData.Has_Milk,
		requestData.Has_Egg,
		requestData.Has_Bean,
		requestData.Has_Seafood,
		requestData.SpcFood,
		requestData.Status,
		requestData.Description,
		ImageID,
	)

	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid ", http.StatusInternalServerError))
		return
	}
	title := "Food"
	description := "create food successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))

}

func updateFood(c *gin.Context) {
	type RequestData struct {
		ID          uint32 `json:"id"`
		Name        string `json:"name"`
		Price       uint32 `json:"price"`
		TypeID      uint32 `json:"type_id"`
		Has_Milk    bool   `json:"has_milk"`
		Has_Egg     bool   `json:"has_egg"`
		Has_Bean    bool   `json:"has_bean"`
		Has_Seafood bool   `json:"has_seafood"`
		Status      bool   `json:"status"`
		Description string `json:"description"`
		ImageID     uint32 `json:"image_id"`
		SpcFood     bool   `json:"spc_food"`
		//image *http.Request `json:"image"`
		NewImg bool `json:"new_img"`
	}
	var requestData RequestData

	if e := c.BindJSON(&requestData); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return

	}
	var ImageID uint32
	if requestData.NewImg {
		ImageID = models.GetIDImage()
	} else {
		ImageID = requestData.ImageID
	}

	err := models.UpdateFood(
		requestData.ID,
		requestData.Name,
		requestData.Price,
		requestData.TypeID,
		requestData.Has_Milk,
		requestData.Has_Egg,
		requestData.Has_Bean,
		requestData.Has_Seafood,
		requestData.SpcFood,
		requestData.Status,
		requestData.Description,
		ImageID,
	)
	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid input ", http.StatusInternalServerError))
		return
	}
	title := "Food"
	description := "update food successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))

}
func updateIngredients(c *gin.Context) {
	var ingredient *models.Ingredients

	if e := c.BindJSON(&ingredient); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	err := models.UpdateIngredients(ingredient)
	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid input ", http.StatusInternalServerError))
		return
	}
	title := "Ingredients"
	description := "update ingredient successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))
}

func deleteFood(c *gin.Context) {
	type RequestData struct {
		ID uint32 `json:"id"`
	}
	var requestData RequestData
	if e := c.BindJSON(&requestData); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	err := models.DeleteFood(
		requestData.ID)

	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("Can not Delete ", http.StatusInternalServerError))
		return
	}
	title := "Delete food"

	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title: &title,
	}))
	return
}
func createingredients(c *gin.Context) {
	type RequestData struct {
		Name   string `json:"name"`
		FoodID uint32 `json:"food_id"`
		Source string `json:"source"`
	}
	var requestData []RequestData
	if e := c.BindJSON(&requestData); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}
	foodid := models.GetFoodID()
	for i := range requestData {
		models.CreateIngredient(
			requestData[i].Name,
			foodid,
			requestData[i].Source,
		)
	}
	title := "Food"
	description := "create ingredient successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))

}

func toggleMenu(c *gin.Context) {
	type RequestData struct {
		ID     uint32 `json:"id"`
		Status bool   `json:"status"`
	}
	var requestData RequestData
	if e := c.BindJSON(&requestData); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}

	err := models.ToggleMenu(
		requestData.ID,
		requestData.Status,
	)
	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid input ", http.StatusInternalServerError))
		return
	}
	title := "Menu"
	description := "Toggle menu successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))
}

func deletePro(c *gin.Context){
	type RequestData struct {
		ID uint32 `json:"id"`
	}
	var requestData RequestData
	if e := c.BindJSON(&requestData); e != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ErrorMessage("Invalid input", http.StatusBadRequest))
		return
	}

	err := models.DeletePro(
		requestData.ID,
	)
	if err != nil {
		c.AbortWithStatusJSON(200, utils.ErrorMessage("invalid input ", http.StatusInternalServerError))
		return
	}
	title := "Promotion"
	description := "Delete promotion successful"
	c.AbortWithStatusJSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
	}))
	
}