package command

import (
	"fmt"

	"github.com/kfngp/gser/framework/cobra"
	"github.com/kfngp/gser/framework/contract"
)

var DemoCommand = &cobra.Command{
	Use:   "demo",
	Short: "demo for framework",
	Run: func(c *cobra.Command, args []string) {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		fmt.Println("app base folder:", appService.BaseFolder())
	},
}
