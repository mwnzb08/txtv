package domain

import "time"

type Test struct {
	Id int64 `xorm:"pk autoincr"`
	Number int64
}

type User struct {
	Id int `xorm:"pk autoincr"`
	Name string
	UserId string
	Pwd string
	CreationDate time.Time `xorm:"created"`
	ModificationDate time.Time `xorm:"updated"`
	Version int `xorm:"version"`
}

