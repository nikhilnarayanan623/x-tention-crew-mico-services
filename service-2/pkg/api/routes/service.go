package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/api/handler/interfaces"
)

func RegisterRoutes(api *gin.RouterGroup, serviceHandler interfaces.ServiceHandler) {

	api.POST("/user", serviceHandler.Method)

}
