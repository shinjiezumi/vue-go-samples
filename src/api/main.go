package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello Go",
		})
	})

	err := router.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
