package middlewares

import (
	"adamnasrudin03/challenge-wallet/pkg/helpers"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var lock = &sync.Mutex{}

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lock.Lock()
		defer lock.Unlock()
		claims, err := helpers.VerifyToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, helpers.APIResponse(err.Error(), http.StatusUnauthorized, nil))
			return
		}

		ctx.Set("userData", claims)
		ctx.Next()
	}
}

func CheckAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		fmt.Println(userData)
		userID := uint64(userData["id"].(float64))
		if userID == 0 {
			return
		}

		ctx.Next()
	}
}
