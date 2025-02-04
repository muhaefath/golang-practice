package controller

import (
	"fmt"
	"golang_practice/request"
	"golang_practice/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

const AuthorizationKey = "ping-123"

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "authorization  is empty"})
			return
		}

		if header != AuthorizationKey {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "wrong authorization"})
			return
		}

		c.Next()

		fmt.Println("Success Authorization")
	}
}

func InitPing(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	group := r.Group("group")
	group.Use(Authorization())
	{
		group.GET("/ping", GetPing)
		group.GET("/ping2", GetPing2)
		group.POST("/ping", PostPing3)
	}
}

func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "group pong",
	})

}

func GetPing2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "group pong 2",
	})
}

func PostPing3(c *gin.Context) {
	var jsonReq request.Ping

	if err := c.ShouldBindJSON(&jsonReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Ping{
		BaseResponse: response.BaseResponse{
			Code:    200,
			Message: fmt.Sprintf("success: %d", jsonReq.ID),
		},
		ID: jsonReq.ID,
	})
}
