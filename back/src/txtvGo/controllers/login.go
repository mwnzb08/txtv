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
	pwd = config.Md5MixEncryption(pwd.(string))
	delete(params, "pwd")
	result := config.SelectDomain(&model, []string{"pwd"}, params)
	if strings.EqualFold(pwd, result) {

	}
	return result
}

