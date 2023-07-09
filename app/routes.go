package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) {
	fmt.Println("[Router] Starting Router")
	router := gin.Default()

	router.GET("/", rootController)
	router.GET("/assets/:file", assetsController)

	router.POST("/messages", func(ctx *gin.Context) {
		sendMessage(db, ctx)
	})
	router.GET("/messages", func(ctx *gin.Context) {
		getMessages(db, ctx)
	})

	router.Run()
}
