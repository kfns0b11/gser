package framework

type NewInstance func(...interface{}) (interface{}, error)

type ServiceProvider interface {
	Boot(Container) error

	Register(Container) NewInstance

	Params(Container) []interface{}

	Name() string

	IsDefer() bool
}
