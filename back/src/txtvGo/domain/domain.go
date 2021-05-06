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
	CreationDate time.Time `xorm:"created"`
	ModificationDate time.Time `xorm:"updated"`
	Version int `xorm:"version"`
}

