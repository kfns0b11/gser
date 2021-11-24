package http

import (
	"github.com/kfngp/gser/app/http/module/demo"
	"github.com/kfngp/gser/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist", "./dist/")
	demo.Register(r) //nolint: errcheck
}
