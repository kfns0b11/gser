package demo

import (
	"fmt"

	"github.com/kfngp/gser/framework"
)

type DemoService struct {
	Service

	c framework.Container
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	cont := params[0]
	c := cont.(framework.Container)
	fmt.Println("new demo service")
	return &DemoService{c: c}, nil
}

func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}
