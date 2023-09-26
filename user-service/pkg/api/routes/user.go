package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/api/handler/interfaces"
)

func RegisterRoutes(api *gin.RouterGroup, userHandler interfaces.UserHandler) {

	api.POST("/user", userHandler.CreateAccount)
	api.GET("/user/:userId", userHandler.GetAccount)
}
