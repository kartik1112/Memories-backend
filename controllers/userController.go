package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/models"
	"github.com/kartik1112/Memories-backend/utils"
)

func Signup(ctx *gin.Context) {

	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON provided",
		})
		return
	}

	hash, err := utils.GenerateHashedPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Password Hashing Error",
		})
		return
	}
	user.Password = hash
	err = user.CreateUser()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully Signed Up.",
	})
}

func Login(ctx *gin.Context) {

	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON provided",
		})
		return
	}

	userPassword := user.Password

	err = user.GetUserByEmail()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email",
		})
		return
	}

	ok := utils.CheckPasswordAndHash(user.Password, userPassword)

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Logged In Failed : Incorrect Password!",
		})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Name, user.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
