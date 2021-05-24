package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"main/config"
)

func init () {
	fmt.Println("----------------> app.go read successful")
}

type AppController struct {
	Session *sessions.Session
	Ctx context.Context
}

var (
	//request  map[string]interface{}
	//render interface{}
)

func (a *AppController) GetMsg () interface{} {
	return config.GetAllMsg()
}
