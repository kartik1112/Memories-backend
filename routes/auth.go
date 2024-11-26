package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/initializers"
	"github.com/kartik1112/Memories-backend/models"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Print(err)
	}
	result := initializers.DB.Create(&user)
	fmt.Print(result.Error)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully Signed Up.",
		"created": result.RowsAffected,
	})
}

func login(ctx *gin.Context) {
	var users []models.User
	result := initializers.DB.Where("email = ? AND password_hash = ?", "df", "nfns").Find(&users)

	if result == nil {
		log.Fatal("No results found")
	}

	for _, user := range users {
		fmt.Printf(user.Name)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Working",
		// "found":   result,
	})
}
