package middlewares

import (
	"adamnasrudin03/challenge-wallet/pkg/helpers"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
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
