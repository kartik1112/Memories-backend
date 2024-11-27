package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/initializers"
	"github.com/kartik1112/Memories-backend/models"
	"golang.org/x/crypto/bcrypt"
)

type SignupBody struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	UserImageUrl string `json:"userimageurl"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(ctx *gin.Context) {

	var body SignupBody

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON provided",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		log.Fatal("Can not Generate Hash : ", err)
		return
	}

	result := initializers.DB.Create(&models.User{
		Email:        body.Email,
		PasswordHash: string(hash),
		Name:         body.Name,
		UserImageUrl: body.UserImageUrl,
	})

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User Already Exist.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully Signed Up.",
		"created": result.RowsAffected,
	})

}

func Login(ctx *gin.Context) {

	var user models.User
	var body LoginBody

	err := ctx.ShouldBindJSON(&body)
	fmt.Print(body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON provided",
		})
		return
	}

	result := initializers.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Logged In Failed : Incorrect Password!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Success!",
		"token":   "snfn",
	})
}
