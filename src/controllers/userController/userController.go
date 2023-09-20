package userController

import (
	"net/http"

	_ "github.com/santos95mat/go-api/docs"
	"github.com/santos95mat/go-api/src/models"
	"github.com/santos95mat/go-api/src/service/userService"

	"github.com/gin-gonic/gin"
)

// CreateUser		godoc
// @Summary			Create user
// @Description		Save user data in DB
// @Param			user body models.UserDto true "Create user"
// @Produce			application/json
// @Tags 			user
// @Sucess			200 {object} user
// @Router			/user [post]
func Create(c *gin.Context) {
	var data models.UserDto
	c.Bind(&data)

	user, err := userService.Create(data)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

// GetAllUsers		godoc
// @Summary			Gett all users
// @Description		Gett all users from DB
// @Produce			application/json
// @Tags 			user
// @Sucess			200 {object} users
// @Router			/user [get]
func GetAll(c *gin.Context) {
	_, ex := c.Get("user")

	if !ex {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	users, err := userService.GetAll()

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// GetUserById		godoc
// @Summary			Get a single user by ID
// @Description		Get a single user from DB
// @Param			id path string true "Find a user by id"
// @Produce			application/json
// @Tags 			user
// @Sucess			200 {object} user
// @Router			/user/{id} [get]
func GetById(c *gin.Context) {
	_, ex := c.Get("user")

	if !ex {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	id := c.Param("id")
	user, err := userService.GetById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// UpdateUserById		godoc
// @Summary				Get a single user by ID
// @Description			Get a single user from DB
// @Param				id path string true "Find a user by id"
// @Param				user body models.UserDto true "Update the find user"
// @Produce				application/json
// @Tags 				user
// @Sucess				200 {object} user
// @Router				/user/{id} [put]
func UpdateById(c *gin.Context) {
	_, ex := c.Get("user")

	if !ex {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	id := c.Param("id")

	var data models.UserDto
	c.Bind(&data)

	user, err := userService.UpdateById(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// DeleteUserById		godoc
// @Summary				Delete a single user by ID
// @Description			Delete a single user from DB
// @Param				id path string true "Find a user by id and delete it"
// @Produce				application/json
// @Tags 				user
// @Sucess				200 {object} "message": "User Deletado com sucesso"
// @Router				/user/{id} [delete]
func DeleteById(c *gin.Context) {
	_, ex := c.Get("user")

	if !ex {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	id := c.Param("id")

	err := userService.DeleteById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Deletado com sucesso",
	})
}
