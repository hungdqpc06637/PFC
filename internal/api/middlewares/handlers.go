package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"web-api/internal/pkg/models/response"

	"github.com/gin-gonic/gin"
)

func NoMethodHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response.FailWithDetailed(ctx, http.StatusMethodNotAllowed, nil, "Method Not Allowed")
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response.FailWithDetailed(ctx, http.StatusNotFound, nil, "The processing function of the request route was not found")
	}
}

func RecoveryHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %v\n", err)
			debug.PrintStack()
			response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, errorToString(err))
			ctx.Abort()
		}
	}()
	ctx.Next()
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
