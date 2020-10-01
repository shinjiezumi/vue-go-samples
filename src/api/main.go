package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shinjiezumi/vue-go-samples/src/api/auth"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/todo"
	"log"
	"net/http"
	"os"
)

func main() {
	database.Initialize()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	router.Use(cors.New(config))

	// Api
	api := router.Group("/api")
	{
		api.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Vue Go Samples",
			})
		})

		api.POST("/login", auth.Login)
		api.GET("/refresh_token", auth.RefreshToken)
		api.POST("/register", auth.Register)

		api.Use(auth.MiddlewareFunc())
		{
			api.GET("/user", auth.CurrentUser)

			api.POST("/todos", todo.Store)
			api.GET("/todos", todo.GetList)
			api.PUT("/todos/:id", todo.Modify)
			api.DELETE("/todos/:id", todo.Remove)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err.Error())
	}
}
