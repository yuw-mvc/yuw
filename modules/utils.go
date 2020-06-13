package modules

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	C "github.com/yuw-mvc/yuw/configs"
	E "github.com/yuw-mvc/yuw/exceptions"
	"math/rand"
	"strings"
	"time"
)

type Utils struct {

}

func NewUtils() *Utils {
	return &Utils {

	}
}

func (util *Utils) SetTimeLocation(toTime string) (*time.Location, error) {
	if toTime == "" {
		toTime = C.LocationAsiaShanghai
	}

	return time.LoadLocation(toTime)
}

func (util *Utils) InterfaceToStruct(data map[interface{}]interface{}, res interface{}) (err error) {
	toMap := util.InterfaceToStringInMap(data)

	if toMap == nil {
		err = E.Err("yuw^mod_util_c", E.ErrPosition())
		return
	}

	err = util.MapToStruct(toMap, res)
	return
}

func (util *Utils) InterfaceToStringInMap(data map[interface{}]interface{}) (toMap map[string]interface{}) {
	if len(data) < 1 {
		return
	}

	toMap = make(map[string]interface{}, len(data))
	for key, val := range data {
		toMap[cast.ToString(key)] = val
	}

	return
}

func (util *Utils) MapToStruct(src interface{}, data interface{}) (err error) {
	strJson, err := json.Marshal(src)
	if err != nil {
		return
	}

	return json.Unmarshal(strJson, &data)
}

func (util *Utils) StructToMap(data interface{}) (res *gin.H, err error) {
	byteData, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = json.Unmarshal(byteData, &res)
	return
}

func (util *Utils) JsonToMap(data []byte) (toMap map[string]interface{}) {
	err := json.Unmarshal(data, &toMap)
	if err != nil {
		return
	}

	return
}

func (util *Utils) StrContains(str string, src ...interface{}) (ok bool, err error) {
	if len(src) < 1 {
		err = E.Err("yuw^mod_util_a", E.ErrPosition())
		return
	}

	for _, val := range src {
		if strings.Contains(str, cast.ToString(val)) {
			ok = true
			return
		}
	}

	return
}

func (util *Utils) IntRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
