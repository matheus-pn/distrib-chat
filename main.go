package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var files embed.FS

func main() {
	router := gin.Default()

	router.GET("/", root)

	router.Run()
}

func root(c *gin.Context) {
	file, err := files.ReadFile("index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	c.Data(200, "text/html", file)
}
