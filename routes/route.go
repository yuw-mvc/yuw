package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	E "github.com/yuw-mvc/yuw/exceptions"
	M "github.com/yuw-mvc/yuw/modules"
	"html/template"
	"strings"
)

type (
	Routes interface {
		Tag() string
		Put(r *gin.Engine, toFunc map[string][]gin.HandlerFunc)
		ToFunc() template.FuncMap
	}

	Rcfg []Routes
	Rtpl []interface{}
	Rarr map[string]map[string][]gin.HandlerFunc

	PoT struct {
		Rcfg *Rcfg
		Rtpl *Rtpl
		Rarr *Rarr
	}
)

const YuwSp string = "->"

var (
	RPoT *PoT
	RMaP map[string]interface{} = map[string]interface{}{}
)

func To(r *gin.Engine) {
	var exc *E.Exceptions = E.NewExceptions()

	/**
	 * Todo: No Routes To Redirect
	**/
	r.NoRoute(exc.NoRoute)

	/**
	 * Todo: No Method To Redirect
	**/
	r.NoMethod(exc.NoMethod)

	for _, to := range *RPoT.Rcfg {
		if _, ok := (*RPoT.Rarr)[to.Tag()]; ok == false {
			continue
		}

		if len((*RPoT.Rarr)[to.Tag()]) == 0 {
			continue
		}

		to.Put(r, (*RPoT.Rarr)[to.Tag()])
	}
}

func ToFunc(tpl ...interface{}) template.FuncMap {
	tplFunc := template.FuncMap{}

	var util *M.Utils = M.NewUtils()
	for _, to := range *RPoT.Rcfg {
		if ok, _ := util.StrContains(to.Tag(), tpl ...); ok == false {
			continue
		}

		if to.ToFunc() == nil {
			continue
		}

		for Tag, toFunc := range to.ToFunc() {
			tplFunc[Tag] = toFunc
		}
	}

	return tplFunc
}

func ToLoggerWithFormatter() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[YuW] %v |	%v |	%v |	%v |	%v |	%v(%v)\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
		)
	})
}

func Do(g *gin.RouterGroup, toFunc map[string][]gin.HandlerFunc) {
	for route, ctrl := range toFunc {
		Y := strings.Split(route, YuwSp)

		if len(Y) != 3 {
			continue
		}

		RMaP[Y[0]] = Y[2]

		switch strings.ToLower(Y[1]) {
		case "get":
			g.GET (Y[2], ctrl ...)
			continue

		case "any":
			g.Any (Y[2], ctrl ...)
			continue

		case "post":
			g.POST(Y[2], ctrl ...)
			continue

		default:
			continue
		}
	}
}
