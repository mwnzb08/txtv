package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12/context"
)
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
	if len(source) == 0 || len(source) < len(valid){
		return false
	}
	// 判断切片的元素是否包含函数需要一个
	//for _,key1 := range valid {
	//	for key2,_ := range source {
	//		if strings.EqualFold(key1,key2) {
	//			return false
	//		}
	//	}
	//}
	return true
}
