package routes

import (
	"github.com/santos95mat/go-blog-api/src/controllers/postController"
	"github.com/santos95mat/go-blog-api/src/middlewares"

	"github.com/gin-gonic/gin"
)

func addPostRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/post")

	posts.POST("/", middlewares.AuthMiddleware, postController.Create)
	posts.GET("/", middlewares.AuthMiddleware, postController.GetAll)
	posts.GET("/:id", middlewares.AuthMiddleware, postController.GetById)
	posts.PUT("/:id", middlewares.AuthMiddleware, postController.UpdateById)
	posts.DELETE("/:id", middlewares.AuthMiddleware, postController.DeleteById)
}
