package controller

import (
	"net/http"

	"backend_gin.com/gin/model"
	"github.com/gin-gonic/gin"
)

func Greet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": "Hello World"})
}


func Name(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{"Result": "OK", "Name": name})
}


func Search(ctx *gin.Context) {
	q := ctx.Query("q")
	page := ctx.Query("page")
	ctx.JSON(http.StatusOK, gin.H{"Result": "OK", "Query": q, "Page": page})
}

func Login(ctx *gin.Context) {
	var user model.LoginRequest

	if err := ctx.ShouldBindJSON(&user); err != nil { // Bind JSON body to struct
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Failed",
			"message": "Invalid request body",
		})
		return
	}
	if user.Username == "Aerospace" && user.Password == "Rockets" {
		ctx.JSON(http.StatusAccepted, gin.H{
			"Result":  "OK",
			"message": "You are authorised",
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Result":  "Failed",
			"message": "You are not authorised",
		})
		return
	}
}
