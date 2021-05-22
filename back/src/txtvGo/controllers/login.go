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
// check 10's when registry user_id if exits
func (c *LoginController) PostCheckRegistryUserId () interface{} {
	c.Session.Increment("countRequest", 1)
	if c.Session.Get("countRequest").(int) > 10 {
		return map[string]string{"gridData": "try again later"}
	}
	request = config.GetJson(c.Ctx)
	render = service.PostCheckRegistryUserId(request)
	return render
}
// registry interface
func (c *LoginController) PostRegistry () interface{} {
	request = config.GetJson(c.Ctx)
	render = service.PostRegistry(request)
	return render
}
// send
func (c *LoginController) PostValidCodeToEmail () {
	request = config.GetJson(c.Ctx)
	service.PostValidCodeToEmail(request)
}

