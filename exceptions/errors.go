package exceptions

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
)

func init() {
	msg = &ErrMsg {
		"yuw^default": "unknown error",
		"yuw^error": "404 error",

		"yuw^ad_a": "modules initialize nil",
		"yuw^ad_b": "Templates Resources nil, Please check the configure",
		"yuw^ad_c": "routes empty",
		"yuw^ad_d": "routes struct empty",
		"yuw^mod_util_a": "utils contains, source is nil",
		"yuw^mod_util_b": "utils merge to map, key exist",
		"yuw^mod_util_c": "utils Interface to Struct, data nil",

		"yuw^m": "error module",
		"yuw^m_b": "error initialize",
		"yuw^m_c": "error env db clusters.databases configures",
		"yuw^m_init_a": "config environment error, go run ... --env=dev|stg|prd",
		"yuw^m_init_b": "config environment error, --env=dev|stg|prd",
		"yuw^m_init_c": "missing .env.dev.yaml",
		"yuw^m_init_d": "config environment, ReadInConfig error",
		"yuw^m_init_e": "config environment, redis configs error",

		"yuw^m_db_a": "error db data source cluster or configs",
		"yuw^m_db_b": "db master error",
		"yuw^m_db_c": "db slaver error",
		"yuw^m_db_d": "table exist error",
		"yuw^m_db_e": "set type in Model PoT struct",

		"yuw^m_rd_a": "error redis engine",
		"yuw^m_rd_b": "redis writer error",
		"yuw^m_rd_c": "redis reader error",

		"yuw^m_fs_a": "open file error",

		"yuw^m_in_a": "lost env languages configs",

		"yuw^m_token_a": "empty key",

		"yuw^m_email_a": "mail configure error",
		"yuw^m_email_b": "mail send configure error",

		"yuw^m_admin_a": "session error",
	}
}

func Err(tag string, content ...interface{}) (err error) {
	var str string = cast.ToString((*msg)["yuw^default"])

	s, ok := (*msg)[tag]
	if ok {
		str = cast.ToString(s)
	}

	if len(content) > 0 {
		str = str + ", " + fmt.Sprint(content ...)
	}

	err = errors.New(str)
	return
}

func ErrArray(arr *ErrType) {
	for tag, ok := range *arr {
		if ok {
			ErrPanic(Err(tag))
		}
	}
}

func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
