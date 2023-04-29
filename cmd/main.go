package main

import (
	"github.com/gin-gonic/gin"
	"lmwn-assignment-go/src"
	"log"
)

const (
	Port = "8080"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("/", src.Health)
	err := router.Run(":" + Port)
	if err != nil {
		log.Println("Failed to run the web server, err:", err)
	}
}