package app

import (
	"bytes"
	"embed"
	"encoding/base64"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//go:embed static/*
var files embed.FS

func rootController(c *gin.Context) {
	file, err := files.ReadFile("static/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error: "+err.Error())
	}

	c.Data(200, "text/html", file)
}

func assetsController(c *gin.Context) {
	f := c.Param("file")
	file, err := files.ReadFile("static/" + f)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error: "+err.Error())
	}

	c.Data(200, "text/html", file)
}

func sendMessage(db *gorm.DB, c *gin.Context) {
	var file *multipart.FileHeader
	var err error

	file, err = c.FormFile("file")

	// If there's no file
	name := c.PostForm("name")
	message := c.PostForm("message")
	if name == "" {
		name = "中国共产党"
	}

	if err != nil {
		db.Create(&Message{
			Username: name, Message: message, MediaUrl: "",
		})

		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	// Access the file data
	fileHeader, err := file.Open()
	if err != nil {
		panic(err)
	}

	defer fileHeader.Close()
	// Read the file data
	fileBytes, err := io.ReadAll(fileHeader)
	if err != nil {
		panic(err)
	}

	lambda := "https://fk5zwqnyjfjws7owbodlzwc42m0huvea.lambda-url.us-east-1.on.aws/"
	// Convert the file data to base64
	fileBase64 := base64.StdEncoding.EncodeToString(fileBytes)

	res, err := http.Post(lambda, "text/plain", bytes.NewBuffer([]byte(fileBase64)))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	jres := gin.H{}
	json.NewDecoder(res.Body).Decode(&jres)

	bucket := "https://distrib-chat-images.s3.amazonaws.com/"
	key := jres["file"].(string)

	db.Create(&Message{
		Username: name, Message: message, MediaUrl: bucket + key,
	})

	c.Redirect(http.StatusSeeOther, "/")
}

func getMessages(db *gorm.DB, c *gin.Context) {
	messages := []Message{}
	db.Find(&messages)
	c.JSON(http.StatusOK, messages)
}
