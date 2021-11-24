package main

import (
	"github.com/kfngp/gser/app/console"
	"github.com/kfngp/gser/app/http"
	"github.com/kfngp/gser/framework"
	"github.com/kfngp/gser/framework/provider/app"
	"github.com/kfngp/gser/framework/provider/kernel"
)

func main() {
	container := framework.NewGserContainer()
	container.Bind(&app.GserAppProvider{}) //nolint: errcheck

	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.GserKernelProvider{HttpEngine: engine}) //nolint: errcheck
	}

	console.RunCommand(container) //nolint: errcheck
}
