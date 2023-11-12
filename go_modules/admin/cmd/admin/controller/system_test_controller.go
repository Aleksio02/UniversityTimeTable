package controller

import (
	stresponse "admin/cmd/admin/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SystemTest(c *gin.Context) {
	systemTestResponse := stresponse.SystemTestResponse{Status: 200, ResponseMessage: "Application is working successfully"}
	c.JSON(http.StatusOK, systemTestResponse)
}
