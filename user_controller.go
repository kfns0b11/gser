package main

import "github.com/kfngp/gser/framework/gin"

func UserLoginConroller(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("ok, UserLoginContronller")
}
