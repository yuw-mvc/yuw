package exceptions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

var (
	EPoT *PoT

	arr *LogMsg
	msg *ErrMsg
	txt *TxTMsg
)

type (
	ErrType map[string]bool

	LogMsg map[string]interface{}
	TxTMsg map[string]interface{}
	ErrMsg map[string]interface{}

	PoT struct {
		ErrMsg *ErrMsg
		TxTMsg *TxTMsg
		LogMsg *LogMsg
	}
)

func PoTCombine() {
	if *EPoT.ErrMsg != nil {
		for key, val := range *EPoT.ErrMsg {
			(*msg)[key] = val
		}
	}

	if *EPoT.TxTMsg != nil {
		for key, val := range *EPoT.TxTMsg {
			(*txt)[key] = val
		}
	}

	if *EPoT.LogMsg != nil {
		for key, val := range *EPoT.LogMsg {
			(*arr)[key] = val
		}
	}
}

func ErrPosition() interface{} {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%v:%v", file, line)
}

type Exceptions struct {

}

func NewExceptions() *Exceptions {
	return &Exceptions {}
}

func (exp *Exceptions) NoRoute(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusNotFound, Err("yuw^error", ErrPosition()))
	return
}

func (exp *Exceptions) NoMethod(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusNotFound, Err("yuw^error", ErrPosition()))
	return
}

