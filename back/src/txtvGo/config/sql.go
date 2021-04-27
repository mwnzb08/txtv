package config

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
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



func SelectDomain(model interface{}, columns []string, params map[string]interface{}) interface{}  {
	fmt.Println("------------------>columns " + strings.Join(columns, ","))
	params2,_ :=json.Marshal(params)
	fmt.Println("------------------>params " +string(params2))
	sqlString := sqlEngine.Cols(strings.Join(columns, ","))
	build := strings.Builder{}
	for key,value := range params {
		build.WriteString(key)
		build.WriteString("=?")
		sqlString.Where( build.String(), value)
		build.Reset()
	}
	if err := sqlString.Unscoped().Find(model); err != nil {
		fmt.Println(">>>>>>>>>>>>Error msg " + err.Error())
	}
	fmt.Println(model)
	return model
}
