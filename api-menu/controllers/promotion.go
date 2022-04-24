package controllers

import (
	"api-menu/models"
	"context"
	"io"
	"net/http"
	"net/url"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func PromotionEndpoint(router *gin.RouterGroup) {
	router.POST("/promotion-image", UploadImagePro)
}

func UploadImagePro(c *gin.Context) {
	ctx := context.Background()
	file, handler, err := c.Request.FormFile("image")
	c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		respondWithJSON(c.Writer, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	imagePath := "promotion/" + handler.Filename

	var bucket = "me-nu-59e93.appspot.com"

	wc1, err := FirebaseStorage.FirebaseStorage.Bucket(bucket)
	if err != nil {
		respondWithJSON(c.Writer, http.StatusBadRequest, err.Error())
		return
	}
	//.Object(imagePath).NewWriter(ctx)
	wc := wc1.Object(imagePath).NewWriter(ctx)
	_, err = io.Copy(wc, file)
	if err != nil {
		respondWithJSON(c.Writer, http.StatusBadRequest, err.Error())
		return

	}
	if err := wc.Close(); err != nil {
		respondWithJSON(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	err = CreateImageProUrl(imagePath, bucket, ctx, FirebaseStorage.client)
	if err != nil {
		respondWithJSON(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(c.Writer, http.StatusCreated, "Create image promotion success.")
}

func CreateImageProUrl(imagePath string, bucket string, ctx context.Context, client *firestore.Client) error {

	type RequestData struct {
		Name string
	}
	var requestData RequestData
	requestData.Name = "https://firebasestorage.googleapis.com/v0/b/" + bucket + "/o/" + url.QueryEscape(imagePath) + "?alt=media"

	models.SaveImageProToDB(
		requestData.Name,
	)

	return nil
}
