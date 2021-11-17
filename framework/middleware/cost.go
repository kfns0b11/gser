package middleware

import (
	"log"
	"time"

	"github.com/kfngp/gser/framework/gin"
)

func Cost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next() //nolint: errcheck

		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", ctx.Request.RequestURI, cost.Seconds())
	}
}
