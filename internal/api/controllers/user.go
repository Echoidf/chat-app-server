package controllers

import (
	"chat-app/internal/api"
	"chat-app/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
)

func GetUsers(c *gin.Context) {
	res, err := api.GetUserList()
	if err != nil {
		log.Println("query user list failed, err:", err.Error())
		response.Error(c, 500, err)
	}

	response.JSON(c, 200, res)
}
