package controller

import (
	"net/http"

	"adamnasrudin03/challenge-wallet/app/dto"
	"adamnasrudin03/challenge-wallet/app/entity"
	"adamnasrudin03/challenge-wallet/app/service"
	"adamnasrudin03/challenge-wallet/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type TransactionController interface {
	Create(ctx *gin.Context)
	TopUp(ctx *gin.Context)
}

type txHandler struct {
	Service *service.Services
}

func NewTransactionController(srv *service.Services) TransactionController {
	return &txHandler{
		Service: srv,
	}
}

func (c *txHandler) Create(ctx *gin.Context) {
	var (
		input dto.CreateTxReq
	)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["id"].(float64))
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

	tx := entity.Transaction{
		UserID:   userID,
		Type:     "OUT",
		Quantity: input.Quantity,
		Name:     input.Name,
		Amount:   input.Amount,
	}

	res, statusHttp, err := c.Service.Transaction.Create(ctx, tx)
	if err != nil {
		ctx.JSON(statusHttp, helpers.APIResponse(err.Error(), statusHttp, nil))
		return
	}

	ctx.JSON(statusHttp, helpers.APIResponse("Success", statusHttp, res))
}

func (c *txHandler) TopUp(ctx *gin.Context) {
	var (
		input dto.TopUpReq
	)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["id"].(float64))
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

	tx := entity.Transaction{
		UserID: userID,
		Type:   "IN",
		Name:   input.BankName,
		Amount: input.Amount,
	}

	res, statusHttp, err := c.Service.Transaction.Create(ctx, tx)
	if err != nil {
		ctx.JSON(statusHttp, helpers.APIResponse(err.Error(), statusHttp, nil))
		return
	}

	ctx.JSON(statusHttp, helpers.APIResponse("Success", statusHttp, res))
}
