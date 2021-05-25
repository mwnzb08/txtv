package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/cache"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"main/controllers"
	"time"
)

func main() {
	app := iris.New()
	//app.Use(iris.Cache304(time.Second * 10))
	sess := sessions.New(sessions.Config{Cookie:"sessionIds", Expires: time.Minute * 10})
	app.Use(sess.Handler())
	sys := mvc.New(app.Party("/"))
	sys.Register(sess.Start)
	sys.Handle(&controllers.LoginController{})
	mvc.Configure(app, cacheController) // 使用緩存的controller
	app.Run(iris.Addr(":8024"))
}

// cacheController base on some request that even no change and less influence
var cacheHandler = cache.Handler(time.Second * 5) // 服务器端缓存10s防止并发
var cache304 = cache.Cache304(time.Second * 10) // 浏览器缓存10s 减少访问
func cacheController (m *mvc.Application) {
	m.Router.Use(cache304, cacheHandler) // 顺序很重要必须先304可能
	m.Handle(&controllers.CacheController{})
}
