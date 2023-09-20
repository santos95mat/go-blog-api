package main

import (
	"github.com/santos95mat/go-api/src/initializers"
	"github.com/santos95mat/go-api/src/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.User{})
}
