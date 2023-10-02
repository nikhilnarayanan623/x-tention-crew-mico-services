package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/api/handler/interfaces"
	usecaseinterface "github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/utils/models"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/utils/models/response"
)

type serviceHandler struct {
	usecase usecaseinterface.UseCase
}

func NewHandler(uc usecaseinterface.UseCase) interfaces.ServiceHandler {
	return &serviceHandler{
		usecase: uc,
	}
}

// @Summary		Method
// @Id				Method
// @Tags			User
// @Param			inputs	body	models.MethodType{}	true	"Details"
// @Router			/user [post]
// @Success		200	{object}	response.Response{data=models.AllUserDetails}	"successfully found all user details"
// @Failure		400	{object}	response.Response{}								"failed to bind input; method:'1 or 2'""
// @Failure		500	{object}	response.Response{}								"failed to get user details"
func (u *serviceHandler) Method(ctx *gin.Context) {

	var body models.MethodType

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse("failed to bind input; method:'1 or 2'", err)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	// method number should be 1 or 2 it's already validation while binding

	var (
		allUser models.AllUserDetails
		err     error
	)

	if body.Method == 1 { // if method is 1 then call the sequentially
		allUser, err = u.usecase.GetSequentially(body.WaitTime)
	} else { // else call the parallel exec
		allUser, err = u.usecase.GetParallel(body.WaitTime)
	}

	if err != nil {
		response := response.ErrorResponse("failed to get user details", err)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.SuccessResponse("successfully found all user details", allUser)
	ctx.JSON(http.StatusOK, response)
}
