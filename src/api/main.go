package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/shinjiezumi/vue-go-samples/src/api/auth"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/searcher"
	"github.com/shinjiezumi/vue-go-samples/src/api/todo_list"
)

func main() {
	if gin.Mode() == gin.DebugMode {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}
	database.Initialize()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	if gin.Mode() == gin.DebugMode {
		config.AllowOrigins = []string{"http://localhost:8080"}
	} else {
		config.AllowOrigins = []string{"https://vgs.shinjiezumi.work"}
	}
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

		searcher.SetupRoute(api)

		api.Use(auth.MiddlewareFunc())
		{
			api.GET("/user", auth.CurrentUser)
			todo_list.SetupRoute(api)
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
