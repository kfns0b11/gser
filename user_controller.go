package main

import (
	"time"

	"github.com/kfngp/gser/framework/gin"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "def")
	time.Sleep(10 * time.Second)
	c.ISetOkStatus().IJson("ok, UserLoginController: " + foo)
}
