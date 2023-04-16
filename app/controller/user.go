package controller

import (
	"net/http"

	"adamnasrudin03/challenge-wallet/app/dto"
	"adamnasrudin03/challenge-wallet/app/service"
	"adamnasrudin03/challenge-wallet/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	MyWallet(ctx *gin.Context)
}

type userController struct {
	Service *service.Services
}

func NewUserController(srv *service.Services) UserController {
	return &userController{
		Service: srv,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var (
		input dto.RegisterReq
	)
	validate := validator.New()
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, nil))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, http.StatusBadRequest, nil))
		return
	}

	userRes, statusHttp, err := c.Service.User.Register(input)
	if err != nil {
		ctx.JSON(statusHttp, helpers.APIResponse(err.Error(), statusHttp, nil))
		return
	}

	ctx.JSON(statusHttp, helpers.APIResponse("Created", statusHttp, userRes))
}

func (c *userController) Login(ctx *gin.Context) {
	var (
		input dto.LoginReq
	)

	validate := validator.New()
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, nil))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, http.StatusBadRequest, nil))
		return
	}

	loginRes, statusHttp, err := c.Service.User.Login(input)
	if err != nil {
		ctx.JSON(statusHttp, helpers.APIResponse(err.Error(), statusHttp, nil))
		return
	}

	ctx.JSON(statusHttp, helpers.APIResponse("Success", statusHttp, loginRes))
}

func (c *userController) MyWallet(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["id"].(float64))

	res, statusHttp, err := c.Service.User.MyWallet(userID)
	if err != nil {
		ctx.JSON(statusHttp, helpers.APIResponse(err.Error(), statusHttp, nil))
		return
	}

	ctx.JSON(statusHttp, helpers.APIResponse("Success", statusHttp, res))
}
