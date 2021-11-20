package demo

const Key = "gser:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
