package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kfngp/gser/framework"
)

func main() {
	core := framework.NewCore()
	RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}

func RegisterRouter(core *framework.Core) {
	core.Get("/user/login", UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}

func SubjectAddController(c *framework.Context) error {
	_ = c.Json(200, "ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	_ = c.Json(200, "ok, SubjectListController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	_ = c.Json(200, "ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	_ = c.Json(200, "ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	_ = c.Json(200, "ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	_ = c.Json(200, "ok, SubjectNameController")
	return nil
}

func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	// mu := sync.Mutex{}
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}
