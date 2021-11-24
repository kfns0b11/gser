package http

import "github.com/kfngp/gser/framework/gin"

func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	Routes(r)
	return r, nil
}
