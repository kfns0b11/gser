package app

import (
	"github.com/kfngp/gser/framework"
	"github.com/kfngp/gser/framework/contract"
)

type GserAppProvider struct {
	BaseFolder string
}

func (p *GserAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewGserApp
}

func (p *GserAppProvider) Boot(container framework.Container) error {
	return nil
}

func (p *GserAppProvider) IsDefer() bool {
	return false
}

func (p *GserAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, p.BaseFolder}
}

func (h *GserAppProvider) Name() string {
	return contract.AppKey
}
