package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8001"

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "Hello World!",
		})
	})

	log.Fatal(router.Run(": " + port))
}
