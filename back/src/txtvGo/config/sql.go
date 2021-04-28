package config

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"reflect"
	"strings"
	"time"
	"xorm.io/core"
)

func init () {
	fmt.Println("----------------> sql.go read successful")
	ConnectDB()
	//if err:=sqlEngine.Sync2(new(tableStruct)); err !=nil {
	//	fmt.Println("mysql table fail")
	//}
}

var sqlEngine *xorm.Engine

func ConnectDB () {
	sqlEngine, _ = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/helloworld?charset=utf8")
	sqlEngine.SetMaxIdleConns(50)
	sqlEngine.SetMaxOpenConns(10)
	sqlEngine.SetMapper(core.GonicMapper{})
	sqlEngine.ShowSQL(true)
	sqlEngine.SetConnMaxLifetime(time.Second * 10)
	if err := sqlEngine.Ping(); err !=nil {
		fmt.Println("mysql connecting fail")
		return
	}
	fmt.Println("mysql connecting successful")
}
// left to right is "entity","column","params"
func SelectDomain(model interface{}, columns []string, params map[string]interface{}) interface{}  {
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
		default: {
			build.WriteString(key)
			build.WriteString("=?")
			sqlString.Where( build.String(), value)
		}
		}
		build.Reset()
	}
	if err := sqlString.Unscoped().Find(model); err != nil {
		fmt.Println(">>>>>>>>>>>>Error msg " + err.Error())
	}
	return model
}
