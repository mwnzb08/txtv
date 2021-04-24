package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
)

type LoginController struct {
	Session *sessions.Session
}

func (c *LoginController) Get() string {
	loginUser := c.Session.Increment("userSession", 1)
	return fmt.Sprintf("%d visit from my current session in seconds of server's up-time",
		loginUser)
}

func (c *LoginController) GetLogin(ctx context.Context) string {
	str := ctx.URLParam("name")
	return str
}

