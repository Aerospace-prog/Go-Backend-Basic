package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "Bearer mysecrettoken" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error":   "Failed",
				"message": "Invalid or missing token",
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"Success": "OK",
				"message": "Token is valid",
			})
		}

		ctx.Next()
	}
}
