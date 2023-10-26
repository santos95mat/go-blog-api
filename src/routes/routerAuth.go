package routes

import (
	"github.com/santos95mat/go-blog-api/src/controllers/authController"

	"github.com/gin-gonic/gin"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/login")

	posts.POST("/", authController.Login)
	posts.POST("/password", authController.ResetPassword)
}
