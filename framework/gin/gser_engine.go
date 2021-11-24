package gin

import "github.com/kfngp/gser/framework"

func (engine *Engine) Bind(p framework.ServiceProvider) error {
	return engine.container.Bind(p)
}

func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container
}
