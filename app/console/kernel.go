package console

import (
	"github.com/kfngp/gser/app/console/command/demo"
	"github.com/kfngp/gser/framework"
	"github.com/kfngp/gser/framework/cobra"
	"github.com/kfngp/gser/framework/command"
)

func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "gser",
		Short: "gser command",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	rootCmd.SetContainer(container)
	command.AddKernelCommands(rootCmd)
	AddAppCommand(rootCmd)

	return rootCmd.Execute()
}

func AddAppCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demo.InitFoo())
}
