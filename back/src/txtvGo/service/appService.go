package service

import (
	"errors"
	"fmt"
	"github.com/kataras/iris/v12/sessions"
	"main/config"
	"main/domain"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func PostLogin (request map[string]interface{}, session *sessions.Session) interface{} {
	if !config.ValidMapKey(request,[]string{"pwd","user_id"}) {
		panic(errors.New("--------->params error"))
	}
	model := make([]domain.User,0)
	pwd := request["pwd"]
	delete(request, "pwd")
	if err := config.SelectDomain(&model, []string{}, request); err != nil {
		fmt.Println("select domain error " + err.Error())
	}
	render := make(map[string]interface{})

	userSession := make(map[string]interface{})
	if len(model) > 0 && strings.EqualFold(config.Md5MixEncryption(pwd.(string)), model[0].Pwd) {
		session.Set("userName", model[0].UserId)
		session.Set("isLogin", true)
		userSession["isLogin"] = true
	} else {
		userSession["isLogin"] = false
		session.Set("isLogin", false)
	}
	render["userSession"] = userSession
	render["gridData"] = model
	return render
}

func PostRegistry (request map[string]interface{}) interface{} {
	if !config.ValidMapKey(request,[]string{"user_id","pwd", "email"}) {
		panic(errors.New("--------->params error"))
	}
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	if len(request["user_id"].(string)) < 6 || len(request["pwd"].(string)) < 6 || !regexp.MustCompile(pattern).MatchString( request["email"].(string)) {
		panic(errors.New("--------->非法操作 error"))
	}
	model := new(domain.User)
	model.UserId = request["user_id"].(string)
	actuallyPwd := config.Md5MixEncryption(request["pwd"].(string))
	model.Pwd = actuallyPwd
	model.Email = request["email"].(string)
	render := make(map[string]interface{})
	if err := config.InsertDomainOne(model); err != nil {
		fmt.Println("insert domain error " + err.Error())
		render["result"] = "appMsg.E0006"
		render["isRegistry"] = false
	} else {
		render["result"] = "appMsg.E0007"
		render["isRegistry"] = true
	}
	return render
}

func PostCheckRegistryUserId (request map[string]interface{}) interface{} {
	if !config.ValidMapKey(request,[]string{"user_id"}) {
		panic(errors.New("--------->params error"))
	}
	model := make([]domain.User,0)
	userId := request["user_id"]
	requestUnique := make(map[string]interface{})
	requestUnique["user_id"] = userId
	render := make(map[string]interface{})
	if err := config.SelectDomain(&model, []string{"user_id"}, requestUnique); err != nil {
		fmt.Println("select domain error " + err.Error())
	}
	if len(model) > 0 {
		render["gridData"] = true
	} else {
		render["gridData"] = false
	}
	return render
}

func PostValidCodeToEmail (request map[string]interface{}, session *sessions.Session) {
	if !config.ValidMapKey(request,[]string{"user_id"}) {
		panic(errors.New("--------->params error"))
	}
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(8999) + 1000
	model := make([]domain.User,0)
	userId := request["user_id"]
	requestUnique := make(map[string]interface{})
	requestUnique["user_id"] = userId
	fmt.Println("sssssssssssssssssssssssss")
	fmt.Println(model)
	if err := config.SelectDomain(&model, []string{"email"}, requestUnique); err != nil {
		fmt.Println("select domain error " + err.Error())
	}
	sendFrom := "system"
	sendTo := []string{model[0].Email}
	body := "<html><body><div>Hello guys you valid code is</div><div>" + strconv.Itoa(randNum) + "</div></body></html>"
	if err:=config.SendEmail(sendFrom, sendTo, "Valid Code",body,"registry"); err !=nil {
		fmt.Println(err)
	}
	session.Set("valid_code_service", strconv.Itoa(randNum))
}
