package config

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/ini.v1"
	"reflect"
	"strings"
	"time"
	"xorm.io/core"
)

func init () {
	fmt.Println("----------------> sql.go read successful")
	ConnectDB()
}

var sqlEngine *xorm.Engine

func ConnectDB () {
	cfg,err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("config.ini Read Error" + err.Error())
	}
	databaseType := cfg.Section("database").Key("name").String()
	dataUrl := cfg.Section("database").Key("url").String()
	fmt.Println(databaseType,dataUrl)
	sqlEngine, _ = xorm.NewEngine(databaseType, dataUrl)
	sqlEngine.SetMaxIdleConns(50)
	sqlEngine.SetMaxOpenConns(10)
	sqlEngine.SetMapper(core.GonicMapper{})
	sqlEngine.ShowSQL(true)
	sqlEngine.SetConnMaxLifetime(time.Second * 30)
	if err := sqlEngine.Ping(); err !=nil {
		fmt.Println("mysql connecting fail")
		return
	}
	fmt.Println("mysql connecting successful")
}
// left to right is "entity","column","params"
func SelectDomain(model interface{}, columns []string, params map[string]interface{}) error {
	fmt.Print("------------------>model ")
	fmt.Println(reflect.TypeOf(model))
	fmt.Println("------------------>columns " + strings.Join(columns, ","))
	params2,_ :=json.Marshal(params)
	fmt.Println("------------------>params " +string(params2))
	sqlString := sqlEngine.Cols(strings.Join(columns, ","))
	build := strings.Builder{}
	for key,value := range params {
		switch key {
		case "sort_desc": sqlString.Desc(value.(string))
		case "sort_asc": sqlString.Asc(value.(string))
		case "limit": sqlString.Limit(int(value.(float64)))
		case "sort": sqlString.OrderBy(value.(string))
		default: {
			build.WriteString(key)
			build.WriteString("=?")
			sqlString.Where( build.String(), value)
		}
		}
		build.Reset()
	}
	return sqlString.Unscoped().Find(model)
}

func InsertDomain(model interface{}) error {
	_,err :=sqlEngine.InsertOne(model)
	return err
}
