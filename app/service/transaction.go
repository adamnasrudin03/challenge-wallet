package service

import (
	"errors"
	"log"
	"net/http"

	"adamnasrudin03/challenge-wallet/app/entity"
	"adamnasrudin03/challenge-wallet/app/repository"

	"github.com/gin-gonic/gin"
)

type TransactionService interface {
	Create(ctx *gin.Context, input entity.Transaction) (res entity.TransactionRes, statusCode int, err error)
}

type transactionSrv struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionService(TransactionRepo repository.TransactionRepository) TransactionService {
	return &transactionSrv{
		TransactionRepository: TransactionRepo,
	}
}

func (srv *transactionSrv) Create(ctx *gin.Context, input entity.Transaction) (res entity.TransactionRes, statusCode int, err error) {
	res, err = srv.TransactionRepository.Create(ctx, input)
	if err == errors.New("your money is not enough") {
		log.Printf("[TransactionService-Create] error create data: %+v \n", err)
		return res, http.StatusBadRequest, err
	}

	if err != nil {
		log.Printf("[TransactionService-Create] error create data: %+v \n", err)
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, err
}
