package domain

import "time"

type Test struct {
	Id int64 `xorm:"pk autoincr"`
	Number int64
}

type User struct {
	Id int `xorm:"pk autoincr"`
	Name string `xorm:"varchar(20) notnull 'name'"`
	UserId string `xorm:"varchar(20) unique notnull 'user_id'"`
	Pwd string `xorm:"varchar(32) notnull 'pwd'"`
	Email string `xorm:"varchar(40) 'email'"`
	CreationDate time.Time `xorm:"created"`
	ModificationDate time.Time `xorm:"updated"`
	Version int `xorm:"version"`
}

type Email struct {
	Id int `xorm:"pk autoincr"`
	UserId string `xorm:"varchar(20) notnull 'user_id'"`
	Type string `xorm:"varchar(10) 'type'"`
	Status string `xorm:"varchar(2) notnull 'status'"`
	SendFrom string `xorm:"varchar(40) notnull 'send_from'"`
	SendTo string `xorm:"varchar(40) notnull 'send_to'"`
	Subject string `xorm:"varchar(100) notnull 'subject'"`
	Body string `xorm:"varchar(2550) 'body'"`
	Remark string `xorm:"varchar(255) 'remark'"`
	CreationDate time.Time `xorm:"created"`
	ModificationDate time.Time `xorm:"updated"`
	Version int `xorm:"version"`
}

