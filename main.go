package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kfngp/gser/framework/gin"
	"github.com/kfngp/gser/framework/middleware"
	"github.com/kfngp/gser/provider/demo"
)

func main() {
	eng := gin.New()
	eng.Bind(&demo.DemoServiceProvider{}) //nolint: errcheck

	eng.Use(gin.Recovery())
	eng.Use(middleware.Cost())

	registerRouter(eng)
	server := &http.Server{
		Handler: eng,
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
}
