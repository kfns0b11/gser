package kernel

import (
	"net/http"

	"github.com/kfngp/gser/framework/gin"
)

type GserKernelService struct {
	engine *gin.Engine
}

func NewGserKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &GserKernelService{
		engine: httpEngine,
	}, nil
}

func (service *GserKernelService) HttpEngine() http.Handler {
	return service.engine
}
