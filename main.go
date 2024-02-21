package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hafiddh/go-gin-auth/db"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error load ENV")
	}

	db.ConnectDB()
}

func main() {

	sqlDB, err := db.DBcon.DB()

	if err != nil {
		log.Println("Error get DB conn")
	}

	defer sqlDB.Close()

	port := os.Getenv("port")

	if port == "" {
		port = "8001"
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "Hello World!",
		})
	})

	log.Fatal(router.Run(": " + port))
}
