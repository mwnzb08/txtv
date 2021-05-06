package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"main/config"
	"main/domain"
	"strings"
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
// user login controller
func (c *LoginController) PostLogin() interface{} {
	params := config.GetJson(c.Ctx) // return map[string]interface{}
	model := make([]domain.User,0)
	pwd := params["pwd"]
	actuallyPwd := config.Md5MixEncryption(pwd.(string))
	delete(params, "pwd")
	if err := config.SelectDomain(&model, []string{}, params); err != nil {
		fmt.Println("select domain error " + err.Error())
	}
	render := make(map[string]interface{})
	userSession := make(map[string]interface{})
	if len(model) > 0 && strings.EqualFold(actuallyPwd, model[0].Pwd) {
		c.Session.Set("userName", model[0].UserId)
		c.Session.Set("isLogin", true)
		userSession["isLogin"] = true
	} else {
		userSession["isLogin"] = false
	}
	render["userSession"] = userSession
	render["gridData"] = model
	return render
}

func (c *LoginController) GetLoginout () {
	c.Session.Clear()
}

func (c *LoginController) PostRegistry () interface{} {
	request :=config.GetJson(c.Ctx)
	model := new(domain.User)
	model.UserId = request["user_id"].(string)
	actuallyPwd := config.Md5MixEncryption(request["pwd"].(string))
	model.Pwd = actuallyPwd
	render := make(map[string]interface{})
	if err := config.InsertDomain(model); err != nil {
		fmt.Println("insert domain error " + err.Error())
		render["result"] = "Account already exist"
	} else {
		render["result"] = "Registry successful"
	}
	return render
}

func (c *LoginController) GetTest () {
	fmt.Println(c.Session.Get("userName"))
}

