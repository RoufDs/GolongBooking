package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.example.com/booking/models"
	"www.example.com/booking/utils"
)

func singup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Couold not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couold not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User authenticated", "token": token})
}
