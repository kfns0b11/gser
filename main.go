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
)

func main() {
	core := gin.New()
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
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

func registerRouter(eg *gin.Engine) {
	eg.GET("/user/login", UserLoginConroller)
}
