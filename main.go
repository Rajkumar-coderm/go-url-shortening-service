package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-short/config"
	"github.com/go-short/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	routes.SetupRoutes(r)

	r.Run(":8080")
}
