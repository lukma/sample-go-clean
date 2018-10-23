package main

import (
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lukma/sample-go-clean/account"
	"github.com/lukma/sample-go-clean/common"
	"github.com/lukma/sample-go-clean/post"
)

func init() {
	env := os.Getenv("ENV")
	filePath := path.Join(os.Getenv("GOPATH"), "/src/github.com/lukma/sample-go-clean/env/"+env+".env")
	if strings.EqualFold(env, "prod") {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := godotenv.Load(filePath); err != nil {
		panic(err.Error())
	}
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	apiGroup := router.Group("/api")
	{
		account.ApplyAccountRoute(apiGroup)
		privateGroup := apiGroup.Group("/")
		privateGroup.Use(common.ApplyJwt)
		{
			post.ApplyContentRoute(privateGroup)
		}
	}
	router.Run(":" + os.Getenv("PORT"))
}
