package userService

import (
	"github.com/santos95mat/go-blog-api/src/entity"
	"github.com/santos95mat/go-blog-api/src/initializers"
	"github.com/santos95mat/go-blog-api/src/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func Create(data models.UserDto) (models.User, error) {
	id := entity.NewID()
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		ID:       id,
		Name:     data.Name,
		LastName: data.LastName,
		Email:    data.Email,
		Password: string(hash),
		Nickname: data.Nickname,
	}
	err = initializers.DB.Create(&user).Error

	return user, err
}

func GetAll() ([]models.User, error) {
	var users []models.User
	err := initializers.DB.Preload(clause.Associations).Find(&users).Error

	return users, err
}

func GetById(id string) (models.User, error) {
	var user models.User
	err := initializers.DB.Preload(clause.Associations).First(&user, "id = ?", id).Error

	return user, err
}

func UpdateById(id string, data models.UserDto) (models.User, error) {
	var user models.User
	user, err := GetById(id)
	if err != nil {
		return user, err
	}
	err = initializers.DB.Model(&user).Updates(
		models.User{
			Name:     data.Name,
			LastName: data.LastName,
			Email:    data.Email,
			Password: data.Password,
			Nickname: data.Nickname,
		},
	).Error

	return user, err
}

func DeleteById(id string) error {
	_, err := GetById(id)
	if err != nil {
		return err
	}

	err = initializers.DB.Delete(&models.User{}, "id = ?", id).Error

	return err
}
