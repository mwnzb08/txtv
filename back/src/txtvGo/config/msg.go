package config

import (
	"fmt"
)

func init () {
	fmt.Println("----------------> msg.go read successful")
}

var msgCodeAndValue = map[string]string {
	"appMsg.E0001": "用户名已存在，或者存在恶意信息",
	"appMsg.E0002": "频繁操作,请稍后再试",
	"appMsg.E0003": "信息输入字段过短",
	"appMsg.E0004": "两次输入密码不一致",
	"appMsg.E0005": "邮箱格式不正确",
	"appMsg.E0006": "验证码错误",
}

func Msg (code string) string {
	return msgCodeAndValue[code]
}

func GetAllMsg () map[string]string{
	return msgCodeAndValue
}
