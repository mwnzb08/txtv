package config

import (
	"fmt"
)

func init () {
	fmt.Println("----------------> msg.go read successful")
}

var msgCodeAndValue = map[string]string {
	"appMsg.E0001": "验证码错误",
	"appMsg.E0002": "频繁操作,请稍后再试",
	"appMsg.E0003": "验证码错误",
	"appMsg.E0004": "验证码错误",
	"appMsg.E0005": "验证码错误",
	"appMsg.E0006": "验证码错误",
}

func Msg (code string) string {
	return msgCodeAndValue[code]
}

func GetAllMsg () map[string]string{
	return msgCodeAndValue
}
