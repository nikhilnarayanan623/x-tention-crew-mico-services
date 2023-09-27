package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateAccount(ctx *gin.Context)
	GetAccount(ctx *gin.Context)
	UpdateAccount(ctx *gin.Context)
	RemoveAccount(ctx *gin.Context)
}
