package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	stresponse "tg-bot/cmd/tg-bot/model/response"
)

func SystemTest(c *gin.Context) {
	systemTestResponse := stresponse.SystemTestResponse{Status: 200, ResponseMessage: "Application is working successfully"}
	c.JSON(http.StatusOK, systemTestResponse)
}
