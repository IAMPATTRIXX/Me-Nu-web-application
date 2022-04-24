package controllers

import (
	"api-menu/models"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"cloud.google.com/go/firestore"
	
	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"

	"google.golang.org/api/option"
)


type Storage struct {
	FirebaseStorage *storage.Client
	FirebaseApp     *firebase.App
	client          *firestore.Client
}

var FirebaseStorage Storage

func InitialFirebaseStorage() {
	//storage main.go
	opt := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	fmt.Println(" firebase.NewApp(context.Background(), nil, opt)", err)
	//storagef ,err := cloud.NewClient(app,opt)

	storageClient, err := app.Storage(context.Background())
	fmt.Println("app.Storage(context.Background())", err)
	FirebaseStorage = Storage{
		FirebaseStorage: storageClient,
		FirebaseApp:     app,
	}
}


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}



func CreateImageUrl(imagePath string, bucket string, ctx context.Context, client *firestore.Client) error {

	type RequestData struct {
		Name string
	}
	var requestData RequestData
	requestData.Name = "https://firebasestorage.googleapis.com/v0/b/" + bucket + "/o/" + url.QueryEscape(imagePath) + "?alt=media"
	// imageStructure := models.Images{
	// 	ImageName: requestData.Name,
	// }
	models.SaveImageToDB(
		requestData.Name,
	)
	

	return nil
}



func (s *Storage) LoadUserImageToFirebase(encode64Image string, fileName string) error {
	onlybase64data := encode64Image[strings.IndexByte(encode64Image, ',')+1:]
	unbased, err := base64.StdEncoding.DecodeString(onlybase64data)
	if err != nil {
		fmt.Println("Cannot decode base64: ", err)
		return err
	}

	imageReader := bytes.NewReader(unbased)
	if imageReader == nil {
		fmt.Println("Unable to read decoded base64 data")
		return err
	}

	ctx := context.Background()
	var buck = "me-nu-59e93.appspot.com"
	bucket, err := s.FirebaseStorage.Bucket(buck)
	if err != nil {
		fmt.Println("s.FirebaseStorage.Bucket: ", err)
		return err
	}

	var resPath = "image/" + fileName

	writer := bucket.Object(resPath).NewWriter(ctx)
	// if _, err = io.Copy(writer, f); err != nil {
	if _, err = io.Copy(writer, imageReader); err != nil {
		fmt.Println("s.FirebaseStorage.Bucket: ", err)
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}
	fmt.Println("Upload to Firebase Path", resPath, "successful")

	CreateImageUrl(resPath, buck, ctx, s.client)
	//fmt.Println(resPath)
	return nil
}

// Delete image
func (s *Storage) DeleteFileFromFilebase(fileName string) error {
	ctx := context.Background()
	bucket, err := s.FirebaseStorage.Bucket("me-nu-59e93.appspot.com")
	if err != nil {
		fmt.Println("s.FirebaseStorage.Bucket: ", err)
		return err
	}

	//var resPath string
	var resPath = "/" + fileName

	err = bucket.Object(resPath).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetImageFromFirebase(fileName string, res *string) error {
	ctx := context.Background()
	resPath := "/" + fileName

	bucket, err := s.FirebaseStorage.Bucket("me-nu-59e93.appspot.com")
	if err != nil {
		fmt.Println("s.FirebaseStorage.Bucket: ", err)
		return err
	}
	// fmt.Println("Getting image form path", resPath, "successful")
	reader, err := bucket.Object(resPath).NewReader(ctx)
	if err != nil {
		fmt.Println("bucket.Object(resPath).NewReader(ctx): ", err, resPath)
		return err
	}

	slurp, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("readFile: unable to read data", err)
		return err
	}

	if len(slurp) > 0 {
		base64Slurp := base64.StdEncoding.EncodeToString(slurp)
		*res = base64Slurp
	}

	return nil
}


