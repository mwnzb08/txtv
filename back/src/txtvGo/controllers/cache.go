package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"main/config"
)

func init () {
	fmt.Println("----------------> cache.go read successful")
}

type CacheController struct {
	Session *sessions.Session
	Ctx context.Context
}

var (
	//request  map[string]interface{}
	//render interface{}
)

func (a *CacheController) GetMsg () interface{} {
	fmt.Println("---------------go new")
	return config.GetAllMsg()
}

