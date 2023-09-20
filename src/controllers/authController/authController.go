package authController

import (
	"net/http"

	_ "github.com/santos95mat/go-api/docs"
	"github.com/santos95mat/go-api/src/models"
	"github.com/santos95mat/go-api/src/service/authService"

	"github.com/gin-gonic/gin"
)

// LoginUser		godoc
// @Summary			Logar user
// @Description		Logar user
// @Param			user body models.UserAuthDto true "Logar user"
// @Produce			application/json
// @Tags 			login
// @Sucess			200 {object} user.Token
// @Router			/login [post]
func Login(c *gin.Context) {
	var data models.UserAuthDto
	c.Bind(&data)

	userToken, err := authService.Login(data)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", userToken.Token, 3600*34*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token":       userToken.Token,
		"currentUser": userToken.User,
	})
}

func ResetPassword(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Aqui vai ficar a rota de resetar a senha",
	})
}
