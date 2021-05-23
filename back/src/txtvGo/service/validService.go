package service

import (
	"errors"
	"github.com/kataras/iris/v12/sessions"
	"main/config"
	"strings"
	"sync"
	"time"
)

var once sync.Once

func CheckLogin (request map[string]interface{}, session *sessions.Session) (map[string]interface{}, error) {
	countLogin := session.Increment("countLogin", 1)
	resultParams := make(map[string]interface{})
	var err error
	resultParams["countLogin"] = countLogin
	if countLogin > 5 {
		validCodeService := session.Get("valid_code_service")
		if validCodeService != nil && strings.EqualFold(request["valid_code_ignore"].(string), validCodeService.(string)) {
			err = nil
		} else {
			resultParams["gridData"] = config.Msg("appMsg.E0001")
			err = errors.New(config.Msg("appMsg.E0001"))
		}
	} else if countLogin > 10 {
		go func() {
			once.Do(func() {
				time.Sleep(time.Minute * 5)
				session.Set("countLogin", 0)
			})
		}()
		resultParams["gridData"] = config.Msg("appMsg.E0002")
		err = errors.New(config.Msg("appMsg.E0002"))
	} else {
		err = nil
	}
	result := make(map[string]interface{})
	result["userSession"] = resultParams
	return result, err
}
