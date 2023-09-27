package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

	userID, err := getParamAsUint(ctx, "userId")
	if err != nil {
		response := response.ErrorResponse("failed to get user id from params as int", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := u.usecase.GetAccount(ctx, userID)
	if err != nil {
		response := response.ErrorResponse("failed to create account", err)
		statusCode := http.StatusInternalServerError
		// check if error is user not exis
		if err == usecase.ErrUserNotExist {
			statusCode = http.StatusNotFound
		}
		ctx.JSON(statusCode, response)
		return
	}

	response := response.SuccessResponse("successfully found user account", user)
	ctx.JSON(http.StatusOK, response)
}
func (u *userHandler) UpdateAccount(ctx *gin.Context) {

	userID, err := getParamAsUint(ctx, "userId")
	if err != nil {
		response := response.ErrorResponse("failed to get user id from params as int", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var body request.User
	// bind input from request body
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse("failed to bind input", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := u.usecase.UpdateAccount(ctx, userID, body)
	if err != nil {
		response := response.ErrorResponse("failed to update user details", err)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.SuccessResponse("successfully user details updated", user)
	ctx.JSON(http.StatusOK, response)
}
func (u *userHandler) RemoveAccount(ctx *gin.Context) {

	userID, err := getParamAsUint(ctx, "userId")
	if err != nil {
		response := response.ErrorResponse("failed to get user id from params as int", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = u.usecase.DeleteUser(ctx, userID)
	if err != nil {
		response := response.ErrorResponse("failed to delete user account", err)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.SuccessResponse("successfully account deleted", nil)
	ctx.JSON(http.StatusOK, response)
}

// get path params as uint from request url
func getParamAsUint(ctx *gin.Context, key string) (uint32, error) {

	param := ctx.Param(key)
	value, err := strconv.ParseUint(param, 10, 32)

	if err != nil || value == 0 {
		return 0, fmt.Errorf("failed to get %s from param as int", key)
	}

	return uint32(value), nil
}
