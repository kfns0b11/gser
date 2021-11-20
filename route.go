package main

import "github.com/kfngp/gser/framework/gin"

func registerRouter(core *gin.Engine) {
	core.GET("/user/login", UserLoginController)

	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}
}
