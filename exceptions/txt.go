package exceptions

import (
	"fmt"
	"github.com/spf13/cast"
)

func init() {
	txt = &TxTMsg {
		"yuw^default":	"unknown error",

		/**
		 * Todo: Log Text
		 */
		"yuw^m_logs_a": "Null",
		"yuw^m_logs_b": "Addr",
		"yuw^m_logs_c": "Type",
		"yuw^m_logs_d": "Header",
		"yuw^m_logs_e": "Status",
		"yuw^m_logs_f": "Time",
		"yuw^m_logs_g": "Error",
	}
}

func TxTErr(tag string, content ... interface{}) (str string) {
	s, ok := (*txt)[tag]
	if ok {
		str = cast.ToString(s)
	} else {
		str = cast.ToString((*txt)["yuw^default"])
	}

	if len(content) > 0 {
		str = str + "," + fmt.Sprint(content ...)
	}

	return
}


