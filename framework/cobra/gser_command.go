package cobra

import (
	"github.com/kfngp/gser/framework"
)

func (comm *Command) SetContainer(container framework.Container) {
	comm.container = container
}

func (comm *Command) GetContainer() framework.Container {
	return comm.Root().container
}
