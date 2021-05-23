package config

import (
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"main/domain"
	"strconv"
	"strings"
	"time"
)

func init () {
	fmt.Println("----------------> sendEmail.go read successful")
}

func SendEmail (sendFrom string, sendTo []string, subject string, body string, types string) error {
	fmt.Println("----------------> send a message to " + strings.Join(sendTo,","))
	if len(strings.Join(sendTo, ",")) == 0 {
		return errors.New("sendTo is Empty")
	}
	data := map[string]string{
		"user": "moqiue0729@163.com",
		"pwd": "AUUNIPCDQDABIORB",
		"host": "smtp.163.com",
		"port": "465",
	}
	getStatus := make(chan string, 1)
	go func () { // 虽然可以忽略的时间，但是还是想用并发
		model := new(domain.Email)
		model.UserId = sendFrom
		model.Type = types
		model.SendFrom = data["user"]
		model.SendTo = strings.Join(sendTo,",")
		model.Subject = subject
		model.Body = body
		select { // 超时5s退出
		case ret:=<-getStatus : model.Status = ret; model.Remark = <-getStatus
		case <-time.After(time.Second * 5) : model.Status = "N"; model.Remark = "time out"
		}
		if err := InsertDomainOne(model); err != nil {
			fmt.Println(err)
		}
	}()

	port,_ := strconv.Atoi(data["port"])
	m :=gomail.NewMessage()
	m.SetHeader("From","Test" + "<" + data["user"] + ">")
	m.SetHeader("To", sendTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(data["host"],port,data["user"],data["pwd"])
	err := d.DialAndSend(m)
	if err != nil {
		getStatus <- "N"
		getStatus <- err.Error()
	} else {
		getStatus <- "Y"
		getStatus <- ""
	}
	return err
}
