package command

import "github.com/kfngp/gser/framework/cobra"

func AddKernelCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(DemoCommand)
	rootCmd.AddCommand(initAppCommand())
}
