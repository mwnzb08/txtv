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
//func (a *CacheController) BeforeActivation(m mvc.BeforeActivation) {
//	m.Handle("GET", "/msg", "GetMsg", iris.Cache304(time.Second * 5), cache.Handler(time.Second * 5))
//} /

func (a *CacheController) GetMsg () interface{} {
	fmt.Println("---------------go new")
	return config.GetAllMsg()
}

