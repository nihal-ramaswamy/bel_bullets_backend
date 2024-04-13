package middleware

import (
	utils "bel_bullets/internal/utils"

	"github.com/gin-gonic/gin"
)

func StravaMiddleWare() gin.HandlerFunc {
	token := utils.GetDotEnvVariable("AUTHORIZATION_TOKEN")

	return func(ctx *gin.Context) {
		ctx.Set("Authorization", "Bearer "+token)
		ctx.Next()
	}
}
