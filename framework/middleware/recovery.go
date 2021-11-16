package middleware

import "github.com/kfngp/gser/framework"

func Recovery() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				ctx.SetStatus(500).Json(err)
			}
		}()

		ctx.Next() //nolint: errcheck
		return nil
	}
}
