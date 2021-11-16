package middleware

import (
	"log"
	"time"

	"github.com/kfngp/gser/framework"
)

func Cost() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		start := time.Now()

		ctx.Next() //nolint: errcheck

		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", ctx.GetRequest().RequestURI, cost.Seconds())

		return nil
	}
}
