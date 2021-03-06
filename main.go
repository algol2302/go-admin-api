package main

import (
	"log"
	"os"

	"github.com/algol2302/go-admin-api/config"
	"github.com/algol2302/go-admin-api/migration"
	"github.com/algol2302/go-admin-api/route"
	"github.com/gin-gonic/gin"
)

func init() {
	db := config.Init()
	migration.Migrate(db)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := route.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
