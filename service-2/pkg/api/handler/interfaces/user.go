package interfaces

import "github.com/gin-gonic/gin"

type ServiceHandler interface {
	Method(ctx *gin.Context)
}
