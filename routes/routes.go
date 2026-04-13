package routes

import (
	"github.com/go-short/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/shorten", controllers.CreateShortURL)
	api.GET("/url/:code", controllers.GetURL)
	api.PUT("/url/:code", controllers.UpdateURL)
	api.DELETE("/url/:code", controllers.DeleteURL)
	api.GET("/url/:code/stats", controllers.GetStats)

	// Redirect route
	r.GET("/r/:code", controllers.RedirectURL)
}
