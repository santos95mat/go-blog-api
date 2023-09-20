package postController

import (
	"net/http"

	_ "github.com/santos95mat/go-api/docs"
	"github.com/santos95mat/go-api/src/models"
	"github.com/santos95mat/go-api/src/service/postService"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatePost		godoc
// @Summary			Create post
// @Description		Save post data in DB
// @Param			post body models.PostDto true "Create post"
// @Produce			application/json
// @Tags 			post
// @Sucess			200 {object} post
// @Router			/post [post]
func Create(c *gin.Context) {
	user, ex := c.Get("user")

	if !ex {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	var id uuid.UUID = user.(models.User).ID

	var data models.PostDto
	c.Bind(&data)

	post, err := postService.Create(data, id)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

// GetAllPosts		godoc
// @Summary			Gett all posts
// @Description		Gett all posts from DB
// @Produce			application/json
// @Tags 			post
// @Sucess			200 {object} posts
// @Router			/post [get]
func GetAll(c *gin.Context) {
	_, ex := c.Get("user")

	if !ex {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	posts, err := postService.GetAll()

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// GetPostById		godoc
// @Summary			Get a single post by ID
// @Description		Get a single post from DB
// @Param			id path string true "Find a post by id"
// @Produce			application/json
// @Tags 			post
// @Sucess			200 {object} post
// @Router			/post/{id} [get]
func GetById(c *gin.Context) {
	_, ex := c.Get("user")

	if !ex {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	id := c.Param("id")
	post, err := postService.GetById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// UpdatePostById		godoc
// @Summary				Get a single post by ID
// @Description			Get a single post from DB
// @Param				id path string true "Find a post by id"
// @Param				post body models.PostDto true "Update the find post"
// @Produce				application/json
// @Tags 				post
// @Sucess				200 {object} post
// @Router				/post/{id} [put]
func UpdateById(c *gin.Context) {
	_, ex := c.Get("user")

	if !ex {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	id := c.Param("id")

	var data models.PostDto
	c.Bind(&data)

	post, err := postService.UpdateById(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// DeletePostById		godoc
// @Summary				Delete a single post by ID
// @Description			Delete a single post from DB
// @Param				id path string true "Find a post by id and delete it"
// @Produce				application/json
// @Tags 				post
// @Sucess				200 {object} "message": "Post Deletado com sucesso"
// @Router				/post/{id} [delete]
func DeleteById(c *gin.Context) {
	_, ex := c.Get("user")

	if !ex {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	id := c.Param("id")

	err := postService.DeleteById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post Deletado com sucesso",
	})
}
