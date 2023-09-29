package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/usecase"
	usecaseinterface "github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/utils/models/request"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/utils/models/response"
)

type userHandler struct {
	usecase usecaseinterface.UserUseCase
}

func NewUserHandler(uc usecaseinterface.UserUseCase) interfaces.UserHandler {
	return &userHandler{
		usecase: uc,
	}
}

// @Summary		Create Account
// @Description	API for create account
// @Id				CreateAccount
// @Tags			User
// @Param			inputs	body	request.User{}	true	"User Details"
// @Router			/user [post]
// @Success		200	{object}	response.Response{data=response.User}	"Successfully logged in"
// @Failure		400	{object}	response.Response{}								"failed to bind input"
// @Failure		409	{object}	response.Response{}								"user already exist"
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

// @Summary		Get Account
// @Description	API for GetAccount
// @Id				GetAccount
// @Tags			User
// @Param			userId	path	int	true	"User ID"
// @Router			/user [get]
// @Success		200	{object}	response.Response{data=response.User}	"successfully found user account"
// @Failure		400	{object}	response.Response{}								"failed to get user id from params as int"
// @Failure		404	{object}	response.Response{}								"user not exist"
// @Failure		500	{object}	response.Response{}								"failed to get account"
func (u *userHandler) GetAccount(ctx *gin.Context) {

	userID, err := getParamAsUint(ctx, "userId")
	if err != nil {
		response := response.ErrorResponse("failed to get user id from params as int", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := u.usecase.GetAccount(ctx, userID)
	if err != nil {
		response := response.ErrorResponse("failed to get account", err)
		statusCode := http.StatusInternalServerError
		// check if error is user not exist
		if err == usecase.ErrUserNotExist {
			statusCode = http.StatusNotFound
		}
		ctx.JSON(statusCode, response)
		return
	}

	response := response.SuccessResponse("successfully found user account", user)
	ctx.JSON(http.StatusOK, response)
}

// @Summary		Update Account
// @Description	API for UpdateAccount
// @Id				UpdateAccount
// @Tags			User
// @Param			userId	path	int	true	"User ID"
// @Param			inputs	body	request.User{}	true	"User Details"
// @Router			/user [put]
// @Success		200	{object}	response.Response{}	"successfully user details updated"
// @Failure		400	{object}	response.Response{}								"failed to bind input"
// @Failure		404	{object}	response.Response{}								"user not exist"
// @Failure		500	{object}	response.Response{}								"failed to update user details"
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
		statusCode := http.StatusInternalServerError
		// check if error is user not exist
		if err == usecase.ErrUserNotExist {
			statusCode = http.StatusNotFound
		}
		ctx.JSON(statusCode, response)
		return
	}

	response := response.SuccessResponse("successfully user details updated", user)
	ctx.JSON(http.StatusOK, response)
}

// @Summary		Remove Account
// @Description	API for RemoveAccount
// @Id				RemoveAccount
// @Tags			User
// @Param			userId	path	int	true	"User ID"
// @Router			/user [delete]
// @Success		200	{object}	response.Response{}	"successfully user details updated"
// @Failure		400	{object}	response.Response{}								"failed to get user id from params as int"
// @Failure		404	{object}	response.Response{}								"user not exist"
// @Failure		500	{object}	response.Response{}								"failed to delete user account"
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
		statusCode := http.StatusInternalServerError
		// check if error is user not exist
		if err == usecase.ErrUserNotExist {
			statusCode = http.StatusNotFound
		}
		ctx.JSON(statusCode, response)
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
