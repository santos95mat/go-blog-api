package main

import (
	_ "github.com/santos95mat/go-blog-api/docs"
	"github.com/santos95mat/go-blog-api/src/initializers"
	"github.com/santos95mat/go-blog-api/src/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

// @title Tag Service API
// @version 1.0
// @description A Tag Service API in Go using Gin

// @host localhost:3000
// @BasePath /
func main() {
	routes.Run()
}
