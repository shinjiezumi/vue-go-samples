package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shinjiezumi/vue-go-samples/src/api/auth"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	// Api
	api := router.Group("/api")
	{
		api.GET("/", auth.HandleRoot)
		api.GET("/login", auth.HandleLogin)
		api.POST("/register", auth.HandleRegister)
	}

	// Static
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	err := router.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
