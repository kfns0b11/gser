package command

import "github.com/kfngp/gser/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(DemoCommand)
	root.AddCommand(initAppCommand())
}
