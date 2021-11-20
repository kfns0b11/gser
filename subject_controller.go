package main

import (
	"fmt"

	"github.com/kfngp/gser/framework/gin"
	"github.com/kfngp/gser/provider/demo"
)

func SubjectListController(c *gin.Context) {
	demoService := c.MustMake(demo.Key).(demo.Service)

	foo := demoService.GetFoo()

	c.ISetOkStatus().IJson(foo)
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController:" + fmt.Sprint(subjectId))

}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
