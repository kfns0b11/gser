package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/kfngp/gser/framework"
	"github.com/kfngp/gser/framework/util"
)

type GserApp struct {
	container  framework.Container
	baseFolder string
}

func (app GserApp) Version() string {
	return "0.0.1"
}

func (app GserApp) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}

	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_ folder parameter, default value is current path")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	return util.GetExecDirectory()
}

func (app GserApp) ConfigFolder() string {
	return filepath.Join(app.BaseFolder(), "config")
}

func (app GserApp) LogFolder() string {
	return filepath.Join(app.StorageFolder(), "log")
}

func (app GserApp) HttpFolder() string {
	return filepath.Join(app.BaseFolder(), "http")
}

func (app GserApp) ConsoleFolder() string {
	return filepath.Join(app.BaseFolder(), "console")
}

func (app GserApp) StorageFolder() string {
	return filepath.Join(app.BaseFolder(), "storage")
}

func (app GserApp) ProviderFolder() string {
	return filepath.Join(app.BaseFolder(), "provider")
}

func (app GserApp) MiddlewareFolder() string {
	return filepath.Join(app.HttpFolder(), "middleware")
}

func (app GserApp) CommandFolder() string {
	return filepath.Join(app.ConsoleFolder(), "command")
}

func (app GserApp) RuntimeFolder() string {
	return filepath.Join(app.StorageFolder(), "runtime")
}

func (app GserApp) TestFolder() string {
	return filepath.Join(app.BaseFolder(), "test")
}

func NewGserApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &GserApp{baseFolder: baseFolder, container: container}, nil
}
