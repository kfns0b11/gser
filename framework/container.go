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

func (cont *GserContainer) PrintProviders() []string {
	cont.lock.RLock()
	defer cont.lock.RUnlock()
	ret := []string{}
	for _, provider := range cont.providers {
		name := provider.Name()
		ret = append(ret, name)
	}
	return ret
}

// #region *GserContainer implement Container interface

func (cont *GserContainer) Bind(provider ServiceProvider) error {
	cont.lock.Lock()
	defer cont.lock.Unlock()

	key := provider.Name()

	cont.providers[key] = provider

	if !provider.IsDefer() {
		if err := provider.Boot(cont); err != nil {
			return err
		}

		params := provider.Params(cont)
		method := provider.Register(cont)
		instance, err := method(params...)
		if err != nil {
			return err
		}
		cont.instances[key] = instance
	}

	return nil
}

func (cont *GserContainer) IsBind(key string) bool {
	return cont.findServiceProvider(key) != nil
}

func (cont *GserContainer) Make(key string) (interface{}, error) {
	return cont.make(key, nil, false)
}

func (cont *GserContainer) MustMake(key string) interface{} {
	ins, err := cont.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return ins
}

func (cont *GserContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return cont.make(key, params, true)
}

// #endregion

func (cont *GserContainer) findServiceProvider(key string) ServiceProvider {
	cont.lock.RLock()
	defer cont.lock.RUnlock()
	if sp, ok := cont.providers[key]; ok {
		return sp
	}
	return nil
}

func (cont *GserContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(cont); err != nil {
		return nil, err
	}

	if params == nil {
		params = sp.Params(cont)
	}
	method := sp.Register(cont)
	ins, err := method(params)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return ins, err
}

func (cont *GserContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	cont.lock.Lock()
	defer cont.lock.Unlock()

	sp := cont.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + "have not register")
	}

	if forceNew {
		return cont.newInstance(sp, params)
	}

	if ins, ok := cont.instances[key]; ok {
		return ins, nil
	}

	ins, err := cont.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}
	cont.instances[key] = ins
	return ins, nil
}
