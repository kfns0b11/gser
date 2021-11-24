package command

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kfngp/gser/framework/cobra"
	"github.com/kfngp/gser/framework/contract"
)

func initAppCommand() *cobra.Command {
	appCommand.AddCommand(appStartCommand)
	return appCommand
}

var appCommand = &cobra.Command{
	Use:   "app",
	Short: "business control command",
	RunE: func(c *cobra.Command, args []string) error {
		c.Help() //nolint: errcheck
		return nil
	},
}

var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "lanuch the web server",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		core := kernelService.HttpEngine()

		server := &http.Server{
			Handler: core,
			Addr:    ":8888",
		}

		go func() {
			server.ListenAndServe() //nolint: errcheck
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}

		return nil
	},
}
