package service

import (
	"errors"
	"log"
	"net/http"

	"adamnasrudin03/challenge-wallet/app/dto"
	"adamnasrudin03/challenge-wallet/app/entity"
	"adamnasrudin03/challenge-wallet/app/repository"
	"adamnasrudin03/challenge-wallet/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	Register(ctx *gin.Context, input dto.RegisterReq) (res entity.User, statusCode int, err error)
	Login(input dto.LoginReq) (res dto.LoginRes, statusCode int, err error)
	MyWallet(userID uint64) (res entity.Wallet, statusCode int, err error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(UserRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: UserRepo,
	}
}

func (srv *userService) Register(ctx *gin.Context, input dto.RegisterReq) (res entity.User, statusCode int, err error) {
	user := entity.User{
		FullName: input.FullName,
		Password: input.Password,
		Email:    input.Email,
	}

	checkUser, _ := srv.userRepository.GetByEmail(user.Email)
	if checkUser.Email != "" {
		err = errors.New("email user has be registered")
		log.Printf("[UserService-Register] error check email: %+v \n", err)
		return res, http.StatusConflict, err
	}

	res, err = srv.userRepository.Register(ctx, user)
	if err != nil {
		log.Printf("[UserService-Register] error create data: %+v \n", err)
		return res, http.StatusInternalServerError, err
	}

	res.Password = ""

	return res, http.StatusCreated, err
}

func (srv *userService) Login(input dto.LoginReq) (res dto.LoginRes, statusCode int, err error) {
	user, err := srv.userRepository.Login(input)
	if errors.Is(err, gorm.ErrRecordNotFound) || user.ID == 0 {
		return res, http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[UserService-Login] error: %+v \n", err)
		return res, http.StatusInternalServerError, err
	}

	res.Token, err = helpers.GenerateToken(user.ID, user.FullName, user.Email)
	if err != nil {
		log.Printf("[UserService-Login] error generate token: %+v \n", err)
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, err
}

func (srv *userService) MyWallet(userID uint64) (res entity.Wallet, statusCode int, err error) {
	res, err = srv.userRepository.MyWallet(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) || res.ID == 0 {
		return res, http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[UserService-MyWallet] error: %+v \n", err)
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, err
}
