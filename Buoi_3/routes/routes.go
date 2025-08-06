package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/controllers"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/api/posts", controllers.GetPosts)
	router.POST("/api/posts", controllers.CreatePost)
	// Thêm các route khác như /api/posts/:id
}