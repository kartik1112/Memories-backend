package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/utils"
)

func CheckAuthentication(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized User",
		})
		return
	}

	userId, err := utils.VerifyJWT(token)

	fmt.Print("userId", userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized User",
		})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
