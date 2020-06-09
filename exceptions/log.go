package exceptions

import (
	"fmt"
	"github.com/spf13/cast"
	"log"
)

func init() {
	arr = &LogMsg {
		"yuw^default": "unknown error",
		"yuw^m_db_a": "error db engine",
	}
}

func LogErr(tag string, content ...interface{}) {
	var str string = cast.ToString((*arr)["yuw^default"])

	s, ok := (*arr)[tag]
	if ok {
		str = cast.ToString(s)
	}

	if len(content) > 0 {
		str = str + "," + fmt.Sprint(content ...)
	}

	log.Println(str)
}

func LogArray(arr *ErrType) {
	for tag, ok := range *arr {
		if ok {
			LogErr(tag)
		}
	}
}
