package kernel

import (
	"github.com/kfngp/gser/framework"
	"github.com/kfngp/gser/framework/contract"
	"github.com/kfngp/gser/framework/gin"
)

type GserKernelProvider struct {
	HttpEngine *gin.Engine
}

func (p *GserKernelProvider) Register(container framework.Container) framework.NewInstance {
	return NewGserKernelService
}

func (p *GserKernelProvider) Boot(container framework.Container) error {
	if p.HttpEngine == nil {
		p.HttpEngine = gin.Default()
	}
	p.HttpEngine.SetContainer(container)
	return nil
}

func (p *GserKernelProvider) IsDefer() bool {
	return false
}

func (p *GserKernelProvider) Params(container framework.Container) []interface{} {
	return []interface{}{p.HttpEngine}
}

func (p *GserKernelProvider) Name() string {
	return contract.KernelKey
}
