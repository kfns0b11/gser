package demo

import (
	demoService "github.com/kfngp/gser/app/provider/demo"
	"github.com/kfngp/gser/framework/gin"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demoService.DemoProvider{}) //nolint: errcheck

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo_post", api.DemoPost)
	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

func (api *DemoApi) Demo(c *gin.Context) {
	//appService := c.MustMake(contract.AppKey).(contract.App)
	//baseFolder := appService.BaseFolder()
	users := api.service.GetUsers()
	usersDTO := UserModelsToUserDTOs(users)
	c.JSON(200, usersDTO)
}

func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	c.JSON(200, usersDTO)
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err) //nolint: errcheck
	}
	c.JSON(200, nil)
}