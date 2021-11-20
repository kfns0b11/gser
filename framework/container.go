package framework

import (
	"errors"
	"sync"
)

type Container interface {
	Bind(provider ServiceProvider) error
	IsBind(key string) bool

	Make(key string) (interface{}, error)
	MustMake(key string) interface{}

	MakeNew(key string, params []interface{}) (interface{}, error)
}

type GserContainer struct {
	Container

	providers map[string]ServiceProvider
	instances map[string]interface{}

	lock sync.RWMutex
}

func NewGserContainer() *GserContainer {
	return &GserContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

func (c *GserContainer) PrintProviders() []string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	ret := []string{}
	for _, provider := range c.providers {
		name := provider.Name()
		ret = append(ret, name)
	}
	return ret
}

// #region *GserContainer implement Container interface

func (c *GserContainer) Bind(provider ServiceProvider) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	key := provider.Name()

	c.providers[key] = provider

	if !provider.IsDefer() {
		if err := provider.Boot(c); err != nil {
			return err
		}

		params := provider.Params(c)
		method := provider.Register(c)
		instance, err := method(params...)
		if err != nil {
			return err
		}
		c.instances[key] = instance
	}

	return nil
}

func (c *GserContainer) IsBind(key string) bool {
	return c.findServiceProvider(key) != nil
}

func (c *GserContainer) Make(key string) (interface{}, error) {
	return c.make(key, nil, false)
}

func (c *GserContainer) MustMake(key string) interface{} {
	ins, err := c.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return ins
}

func (c *GserContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return c.make(key, params, true)
}

// #endregion

func (c *GserContainer) findServiceProvider(key string) ServiceProvider {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if sp, ok := c.providers[key]; ok {
		return sp
	}
	return nil
}

func (c *GserContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(c); err != nil {
		return nil, err
	}

	if len(params) == 0 {
		params = sp.Params(c)
	}
	method := sp.Register(c)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return ins, err
}

func (c *GserContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	sp := c.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + "have not register")
	}

	if forceNew {
		return c.newInstance(sp, params)
	}

	if ins, ok := c.instances[key]; ok {
		return ins, nil
	}

	ins, err := c.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}
	c.instances[key] = ins
	return ins, nil
}
