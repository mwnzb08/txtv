package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"main/config"
	"main/domain"
)

func init () {
	fmt.Println("----------------> login.go read successful")
}

type LoginController struct {
	Session *sessions.Session
	Ctx context.Context
}

func (c *LoginController) Get() string {
	loginUser := c.Session.Increment("userSession", 1)
	return fmt.Sprintf("%d visit from my current session in seconds of server's up-time",
		loginUser)
}

func (c *LoginController) PostLogin() interface{} {
	model := make([]domain.User,0)
	param := &domain.User{}
	params := make(map[string]interface{})
	if err :=c.Ctx.ReadJSON(param); err != nil {
		fmt.Println("ReadJson has error " + err.Error())
	}
	params["name"] = param.Name
	params["pwd"] = param.Pwd
	result := config.SelectDomain(&model, []string{"dpid,pwd"}, params)
	fmt.Println(result)
	return result
}

