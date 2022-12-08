package controllers

import (
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.ConnectDB("go_rest_api")
}

func GetUsers(ctx *gin.Context) {
	users := models.GetUsers(utils.DB)

	ctx.JSON(http.StatusOK, users)
}

func PostUsers(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	models.SetUser(utils.DB, user.Username, user.Password, user.Balance)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully insert the data",
	})
}

func PutUsers(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	password := "s4n94tr4h4514"
	balance := 0
	models.UpdateUser(utils.DB, id, password, balance)

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func DeleteUsers(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	models.DeleteUser(utils.DB, id)

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
