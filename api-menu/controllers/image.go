package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ImageFirebaseEndpoint(router *gin.RouterGroup) {
	router.POST("upload-image", UploadImage)
	// router.GET("get", GetImagefromFirbase)

}

func UploadImageToFirebase(c *gin.Context) {
	file, handler, err := c.Request.FormFile("image")

	if err != nil {
		log.Fatal("errrr -->", err)
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return

	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)

	}
	
	//var res *string
	imgBase64Str := base64.StdEncoding.EncodeToString(data)
	//fmt.Println(imgBase64Str)
	//FirebaseStorage.CreateImageUrl(imgBase64Str)
	FirebaseStorage.LoadUserImageToFirebase(imgBase64Str, handler.Filename)
	//FirebaseStorage.GetImageFromFirebase(handler.Filename,res)
}



func UploadImage(c *gin.Context) {
	ctx := context.Background()
	file, handler, err := c.Request.FormFile("image")
	c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		respondWithJSON(c.Writer, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	imagePath := "image/" + handler.Filename

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

	err = CreateImageUrl(imagePath, bucket, ctx, FirebaseStorage.client)
	if err != nil {
		respondWithJSON(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(c.Writer, http.StatusCreated, "Create image success.")
}
