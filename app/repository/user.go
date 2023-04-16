package repository

import (
	"errors"
	"log"

	"adamnasrudin03/challenge-wallet/app/dto"
	"adamnasrudin03/challenge-wallet/app/entity"
	"adamnasrudin03/challenge-wallet/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Register(ctx *gin.Context, input entity.User) (res entity.User, err error)
	Login(input dto.LoginReq) (res entity.User, er error)
	GetByEmail(email string) (res entity.User, err error)
	MyWallet(userID uint64) (res entity.Wallet, err error)
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		DB: db,
	}
}

func (repo *userRepo) Register(ctx *gin.Context, input entity.User) (res entity.User, err error) {
	query := repo.DB.Begin().WithContext(ctx)
	myWallet := entity.Wallet{}
	err = query.Create(&input).Error
	if err != nil {
		query.Rollback()
		log.Printf("[UserRepository-Register] error register new user: %+v \n", err)
		return input, err
	}

	myWallet.UserID = input.ID
	err = query.Create(&myWallet).Error
	if err != nil {
		query.Rollback()
		log.Printf("[UserRepository-Register] error create my wallet : %+v \n", err)
		return input, err
	}

	err = query.Commit().Error
	if err != nil {
		query.Rollback()
		log.Printf("[UserRepository-Register] error commit tx: %+v \n", err)
		return
	}
	return input, err
}

func (repo *userRepo) Login(input dto.LoginReq) (res entity.User, err error) {
	if err = repo.DB.Where("email = ?", input.Email).Take(&res).Error; err != nil {
		log.Printf("[UserRepository-Login] error login: %+v \n", err)
		return
	}

	if !helpers.PasswordValid(res.Password, input.Password) {
		err = errors.New("invalid password")
		log.Printf("[UserRepository-Login] error cek pass: %+v \n", err)
		return
	}
	return
}

func (repo *userRepo) GetByEmail(email string) (res entity.User, err error) {
	if err = repo.DB.Where("email = ?", email).Take(&res).Error; err != nil {
		log.Printf("[UserRepository-GetByEmail] error : %+v \n", err)
		return
	}
	return
}

func (repo *userRepo) MyWallet(userID uint64) (res entity.Wallet, err error) {
	if err = repo.DB.Preload(clause.Associations).Where("user_id = ?", userID).Take(&res).Error; err != nil {
		log.Printf("[UserRepository-MyWallet] error : %+v \n", err)
		return
	}
	return
}
