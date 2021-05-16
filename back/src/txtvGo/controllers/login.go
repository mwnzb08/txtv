package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"main/config"
	"main/service"
)

func init () {
	fmt.Println("----------------> login.go read successful")
}

type LoginController struct {
	Session *sessions.Session
	Ctx context.Context
}

var (
	request  map[string]interface{}
	render interface{}
)

func (c *LoginController) Get() string {
	loginUser := c.Session.Increment("userSession", 1)
	return fmt.Sprintf("%d visit from my current session in seconds of server's up-time",
		loginUser)
}
// user login controller
func (c *LoginController) PostLogin() interface{} {
	request = config.GetJson(c.Ctx) // return map[string]interface{}
	render = service.PostLogin(request, c.Session)
	return render
}
// user login out
func (c *LoginController) GetLoginOut () interface{} {
	c.Session.Clear()
	render = map[string]bool{"gridData":true}
	return render
}

func (c *LoginController) PostCheckRegistryUserId () interface{} {
	request = config.GetJson(c.Ctx)
	render = service.PostCheckRegistryUserId(request)
	return render
}

func (c *LoginController) PostRegistry () interface{} {
	request = config.GetJson(c.Ctx)
	render = service.PostRegistry(request)
	return render
}

func (c *LoginController) GetTest () {
	fmt.Println(c.Session.Get("userName"))
}

func (c *LoginController) PostValidCodeToEmail () {
	request = config.GetJson(c.Ctx)
	service.PostValidCodeToEmail(request)
}

