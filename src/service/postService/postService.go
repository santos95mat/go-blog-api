package postService

import (
	"github.com/santos95mat/go-api/src/entity"
	"github.com/santos95mat/go-api/src/initializers"
	"github.com/santos95mat/go-api/src/models"

	"github.com/google/uuid"
)

func Create(data models.PostDto, userID uuid.UUID) (models.Post, error) {
	id := entity.NewID()
	post := models.Post{
		Title:  data.Title,
		Body:   data.Body,
		ID:     id,
		UserID: userID,
	}
	err := initializers.DB.Create(&post).Error

	return post, err
}

func GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := initializers.DB.Find(&posts).Error

	return posts, err
}

func GetById(id string) (models.Post, error) {
	var post models.Post
	err := initializers.DB.First(&post, "id = ?", id).Error

	return post, err
}

func UpdateById(id string, data models.PostDto) (models.Post, error) {
	var post models.Post
	post, err := GetById(id)
	if err != nil {
		return post, err
	}
	err = initializers.DB.Model(&post).Updates(
		models.Post{
			Title: data.Title,
			Body:  data.Body,
		},
	).Error

	return post, err
}

func DeleteById(id string) error {
	_, err := GetById(id)
	if err != nil {
		return err
	}

	err = initializers.DB.Delete(&models.Post{}, "id = ?", id).Error

	return err
}
