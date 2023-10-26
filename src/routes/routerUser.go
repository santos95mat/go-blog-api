package routes

import (
	"github.com/santos95mat/go-blog-api/src/controllers/userController"
	"github.com/santos95mat/go-blog-api/src/middlewares"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/user")

	posts.POST("/", userController.Create)
	posts.GET("/", middlewares.AuthMiddleware, userController.GetAll)
	posts.GET("/:id", middlewares.AuthMiddleware, userController.GetById)
	posts.PUT("/:id", middlewares.AuthMiddleware, userController.UpdateById)
	posts.DELETE("/:id", middlewares.AuthMiddleware, userController.DeleteById)
}
