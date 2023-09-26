package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/usecase"
	usecaseinterface "github.com/nikhilnarayanan623/x-tention-crew/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils/models/request"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils/models/response"
)

type userHandler struct {
	usecase usecaseinterface.UserUseCase
}

func NewUserHandler(uc usecaseinterface.UserUseCase) interfaces.UserHandler {
	return &userHandler{
		usecase: uc,
	}
}

func (u *userHandler) CreateAccount(ctx *gin.Context) {

	var body request.User
	// bind input from request body
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse("failed to bind input", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	// create user account
	user, err := u.usecase.CreateAccount(ctx, body)
	if err != nil {
		response := response.ErrorResponse("failed to create account", err)
		statusCode := http.StatusInternalServerError
		// if error is user already exist change status code to status conflict
		if err == usecase.ErrUserAlreadyExist {
			statusCode = http.StatusConflict
		}
		ctx.JSON(statusCode, response)
		return
	}

	response := response.SuccessResponse("successfully account created", user)

	ctx.JSON(http.StatusCreated, response)
}
func (u *userHandler) GetAccount(ctx *gin.Context) {

}
func (u *userHandler) UpdateAccount(ctx *gin.Context) {

}
func (u *userHandler) RemoveAccount(ctx *gin.Context) {

}
