package auth

import (
	"net/http"
	auth_service "github.com/tests/internal/auth"
	"github.com/tests/internal/service/hash"
	"github.com/tests/internal/service/request"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	user User
	auth Auth
}

func NewController(user User, auth Auth) *Controller {
	return &Controller{user, auth}
}

func (ac Controller) SignIn(c *gin.Context) {
	var data auth_service.SignIn

	if er := request.BindFunc(c, &data, "FirstName", "Password"); er != nil {
		c.JSON(er.Status, gin.H{
			"message": er.Err.Error(),
			"status":  false,
		})

		return
	}

	userDetail, err := ac.user.GetByFirstName(c, data.FirstName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "user not found",
			"status":  false,
		})
		return
	}
	answer := hash.CheckPasswordHash(data.Password, *userDetail.Password)
	if answer == false {
		c.JSON(http.StatusOK, gin.H{
			"message": "incorrect password!",
			"status":  false,
		})
		return
	}

	var generateTokenData auth_service.GenerateToken

	generateTokenData.FirstName = userDetail.FirstName
	generateTokenData.Username = userDetail.Username
	token, err2 := ac.auth.GenerateToken(c, generateTokenData)

	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err2.Error(),
			"status":  false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
