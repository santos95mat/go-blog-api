package authService

import (
	"os"
	"time"

	"github.com/santos95mat/go-blog-api/src/initializers"
	"github.com/santos95mat/go-blog-api/src/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(data models.UserAuthDto) (models.UserToken, error) {
	var user models.User

	err := initializers.DB.First(&user, "email = ?", data.Email).Error

	if err != nil {
		return models.UserToken{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return models.UserToken{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	rokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return models.UserToken{}, err
	}

	userToken := models.UserToken{
		User:  user,
		Token: rokenString,
	}

	return userToken, err
}
