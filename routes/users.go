package routes

import (
	"example/restapi/models"
	"example/restapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the request header",
		})
		return
	}
	err = user.Validate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong password",
		})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong password",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})

}

// signup
func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the request header",
		})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not store the user data",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})

}
