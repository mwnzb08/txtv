package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"main/controllers"
)

func main() {
	app := iris.New()
	sess := sessions.New(sessions.Config{Cookie:"sessionId"})
	app.Use(sess.Handler())
	sys := mvc.New(app.Party("/"))
	sys.Register(sess.Start)
	sys.Handle(&controllers.LoginController{})
	app.Run(iris.Addr(":8081"))
}
