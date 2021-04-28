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
	var resultJson interface{}
	if err := json.Unmarshal(buffer.Bytes() ,&resultJson); err !=nil {
		fmt.Println("json Unmarshal fail " + err.Error())
	}
	return resultJson.(map[string]interface{})
}
