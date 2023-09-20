package routes

import (
	"github.com/santos95mat/go-api/src/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r = gin.Default()

func Run() {
	r.Use(cors.Default())

	//add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	getRoutes()
	r.Run()
}

func getRoutes() {
	v1 := r.Group("")
	addPostRoutes(v1)
	addUserRoutes(v1)
	addAuthRoutes(v1)
}
