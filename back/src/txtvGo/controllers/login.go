package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"main/config"
	"main/service"
	"sync"
	"time"
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

var once sync.Once

// user login controller
func (c *LoginController) PostLogin() interface{} {
	request = config.GetJson(c.Ctx) // return map[string]interface{}
	if resultOfRedis, err := config.Redis.Get(request["user_id"].(string)).Result(); err != nil { // if redis not exists this user cache
		if result, err := service.CheckLogin(request, c.Session); err != nil {
			return result
		}
		render = service.PostLogin(request, c.Session)
	} else {
		return resultOfRedis
	}
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
		go func() {
			once.Do(func() {
				time.Sleep(time.Minute * 5)
				c.Session.Set("countRequest", 0)
			})
		}()
		return map[string]string{"gridData": "try again later"}
	}
	request = config.GetJson(c.Ctx)
	render = service.PostCheckRegistryUserId(request)
	return render
}
// registry interface 没有对频繁请求做后端的限制，也就是可以被ddos攻击。懒
func (c *LoginController) PostRegistry () interface{} {
	request = config.GetJson(c.Ctx)
	render = service.PostRegistry(request)
	return render
}
// send
func (c *LoginController) PostValidCodeToEmail () {
	request = config.GetJson(c.Ctx)
	service.PostValidCodeToEmail(request, c.Session)
}

