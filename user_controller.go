package main

import "github.com/kfngp/gser/framework"

func UserLoginConroller(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("ok, UserLoginContronller")
	return nil
}
