package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12/context"
	"strings"
)

func init () {
	fmt.Println("----------------> getJson.go read successful")
}

// get context data then Unmarshal as json
func GetJson(ctx context.Context) map[string]interface{} {
	var (
		data []byte
		err error
	)
	if data,err = ctx.GetBody(); err != nil {
		fmt.Println("json GetBody fail " + err.Error())
	}
	var buffer bytes.Buffer
	if err := json.Indent(&buffer, data, "", ""); err != nil {
		fmt.Println("json Indent fail " + err.Error())
	}
	var resultJson map[string]interface{}
	if err := json.Unmarshal(buffer.Bytes() ,&resultJson); err !=nil {
		fmt.Println("json Unmarshal fail " + err.Error())
	}
	return resultJson
}

func ValidMapKey(source map[string]interface{}, valid []string) bool {
	fmt.Println(source)
	if len(source) == 0 || len(source) < len(valid){
		return false
	}
	var count int = 0 // 判断验证成功的个数
	for _,key1 := range valid {
		for key2,_ := range source {
			if strings.EqualFold(key1,key2) {
				count ++
			}
		}
	}
	if count != len(valid) {
		return false
	}
	return true
}
