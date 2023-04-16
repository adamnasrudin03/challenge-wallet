package repository

import (
	"errors"
	"log"

	"adamnasrudin03/challenge-wallet/app/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx *gin.Context, input entity.Transaction) (res entity.TransactionRes, err error)
}

type txRepo struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &txRepo{
		DB: db,
	}
}

func (repo *txRepo) Create(ctx *gin.Context, input entity.Transaction) (res entity.TransactionRes, err error) {
	query := repo.DB.Begin().WithContext(ctx)
	myWallet := entity.Wallet{}

	if err = query.Where("user_id = ?", input.UserID).Take(&myWallet).Error; err != nil {
		log.Printf("[TransactionRepository-Create] error get wallet by user id %v : %+v \n", input.UserID, err)
		return
	}

	if myWallet.Amount < input.Amount && input.Type == "OUT" {
		err = errors.New("your money is not enough")
		return
	}

	err = query.Create(&input).Error
	if err != nil {
		query.Rollback()
		log.Printf("[TransactionRepository-Create] error Create new Transaction: %+v \n", err)
		return
	}

	if input.Type == "IN" {
		myWallet.Amount = myWallet.Amount + input.Amount
	} else if input.Type == "OUT" {
		myWallet.Amount = myWallet.Amount - input.Amount
	}

	err = query.Model(entity.Wallet{}).Where("user_id = ?", input.UserID).Updates(&myWallet).Error
	if err != nil {
		query.Rollback()
		log.Printf("[TransactionRepository-Create] error Update wallet: %+v \n", err)
		return
	}

	err = query.Commit().Error
	if err != nil {
		query.Rollback()
		log.Printf("[TransactionRepository-Create] error commit tx: %+v \n", err)
		return
	}

	res = entity.TransactionRes{
		ID:        input.ID,
		UserID:    input.UserID,
		Amount:    input.Amount,
		Quantity:  input.Quantity,
		Name:      input.Name,
		Type:      input.Type,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}

	return res, err
}
