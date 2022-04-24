package utils

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Logging(c *gin.Context, path string, functionName string) {
}

func GetBodyJSON(c *gin.Context, data map[string]interface{}, serviceName string, method string, paramName []string) []byte {
	var err error

	switch method {
	case "GET":
		for index := 0; index < len(paramName); index++ {
			paramValue := c.Query(paramName[index])
			if paramValue != "" {
				data[paramName[index]], err = strconv.Atoi(paramValue)
				if err != nil {
					data[paramName[index]] = paramValue
				}
			}
		}

	case "POST":
		c.ShouldBindWith(&data, binding.JSON)
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
	}

	return dataJSON
}

const (
	AuthErrorMessage = ""
	// AuthRequiredError is a message sending when there is no Authorization field in the header
	AuthRequiredError = "authorization required"
	TokenError        = "invalid token"
	PermissionError   = "permission denied"
)

// ErrorMessagePrototype - a prototype for error message
type ErrorMessagePrototype struct {
	Error errorObject `json:"error"`
}

// SuccessMessagePrototype -- a prototype for success message
type SuccessMessagePrototype struct {
	Data DataObject `json:"data"`
}

type errorObject struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ResponseMessageObject struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type DataObject struct {
	Kind        *string     `json:"kind"`
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	Status      *string     `json:"status"`
	Item        interface{} `json:"item"`
	Items       interface{} `json:"items"`
}
type Timestamp struct {
	time.Time
}

const RFC3339Millis = "2006-01-02T15:04:05.000Z07:00"

// ErrorMessage - return an error message
func ErrorMessage(message string, code int) ErrorMessagePrototype {
	err := errorObject{
		Code:    code,
		Message: message,
	}

	return ErrorMessagePrototype{Error: err}
}

// SuccessMessage - return an success message
func SuccessMessage(data DataObject) SuccessMessagePrototype {
	return SuccessMessagePrototype{Data: data}
}

func StatusOK(c *gin.Context, serviceName string, method string) {
	title := serviceName
	description := "a admin is successfully " + method
	c.AbortWithStatusJSON(http.StatusOK, SuccessMessage(DataObject{
		Title:       &title,
		Description: &description,
	}))
}

func StatusInternalServerError(c *gin.Context, serviceName string, method string) {
	c.AbortWithStatusJSON(200, ErrorMessage("failed to "+method+" "+serviceName, 500))
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	if x == "0001-01-01T00:00:00Z" {
		return true
	}
	if x == nil {
		return true
	}
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

// convert - change bytes array to string
func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, "")
}

func ConvertStringToUInt32(src string) uint32 {
	resUint64, err := strconv.ParseUint(src, 10, 32)
	var res uint32
	if err == nil {
		res = uint32(resUint64)
	}
	return res
}
