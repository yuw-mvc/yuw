package yuw

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	E "github.com/yuw-mvc/yuw/exceptions"
	M "github.com/yuw-mvc/yuw/modules"
	R "github.com/yuw-mvc/yuw/routes"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	defaultPost string = "8888"
	defaultHTMLDir string = "./resources/templates/"
)

func init() {
	E.ErrArray(&E.ErrType{"yuw^ad_a":M.I == nil})
}

type Engine struct {
	PoTRoutes *R.PoT
	PoTExceptions *E.PoT
}

func New() *Engine {
	return &Engine {}
}

func (yuw *Engine) Run() {
	r := yuw.ginInitialized()
	R.To(r)

	/**
	 * Todo: Loading Templates
	**/
	bTpl := cast.ToBool(M.I.Get("Temp.Status", false))
	if bTpl {
		strResources := cast.ToString(M.I.Get("Temp.Resources", ""))
		E.ErrArray(&E.ErrType{"yuw^ad_b":strResources == ""})

		strDirViewer := cast.ToString(M.I.Get("Temp.DirViews", defaultHTMLDir + "viewer/"))
		arrResources := strings.Split(strResources, ":")

		objTPL := multitemplate.NewRenderer()
		for _, skeleton := range arrResources {
			views, _ := ioutil.ReadDir(strDirViewer + skeleton)
			for _, view := range views {
				arrTPL := yuw.tplLoading(skeleton, view.Name())

				/* Todo: need to update */
				if yuw.PoTRoutes.Rtpl != nil {
					objTPL.AddFromFilesFuncs(view.Name(), R.ToFunc(yuw.PoTRoutes.Rtpl), arrTPL ...)
				}
			}
		}
		r.HTMLRender = objTPL
	}

	r.Run(":" + cast.ToString(M.I.Get("Yuw.Port", defaultPost)))
}

func (yuw *Engine) YuwInitialized() *Engine {
	E.ErrArray(&E.ErrType {
		"yuw^ad_c": yuw.PoTRoutes.Rarr == nil,
		"yuw^ad_d": yuw.PoTRoutes.Rcfg == nil,
	})

	R.RPoT = yuw.PoTRoutes
	E.EPoT = yuw.PoTExceptions
	E.PoTCombine()

	return yuw
}

func (yuw *Engine) ginInitialized() (r *gin.Engine) {
	/**
	 * Todo: Log in Command
	 */
	gin.ForceConsoleColor()

	env := cast.ToString(M.I.Env.Get("env"))
	if env == "prd" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	/**
	 * Todo: Start Routes (gin.New() | gin.Default())
	 */
	r = gin.New()
	r.Use(gin.Recovery())

	if env == "prd" {
		r.Use(R.ToLoggerWithFormatter())
	} else {
		r.Use(gin.Logger())
	}

	/**
	 * Todo: Configure Static Resources
	 */
	TempStaticStatus := cast.ToBool(M.I.Get("Temp.StaticStatus", false))
	if TempStaticStatus {
		static := cast.ToString(M.I.Get("Temp.Static", "./resources/assets"))
		staticIcon := cast.ToString(M.I.Get("Temp.StaticIcon", "./resources/favicon.ico"))

		r.Static("./assets", static)
		r.StaticFile("./favicon.ico", staticIcon)
	}

	return
}

func (yuw *Engine) tplLoading(skeleton string, view string) (arrTPL []string) {
	TplSuffix := cast.ToString(M.I.Get("Temp.Suffix", "html"))
	dirLayout := cast.ToString(M.I.Get("Temp.DirLayout", defaultHTMLDir + "layout/"))

	TplLayout, errLayout := filepath.Glob(dirLayout + skeleton + "." + TplSuffix)
	E.ErrPanic(errLayout)

	TplShared := cast.ToString(M.I.Get("Temp.DirShared", defaultHTMLDir + "shared/"))

	shareds, errShared := filepath.Glob(TplShared + skeleton + "/" + "*.html")
	E.ErrPanic(errShared)

	TplViews := cast.ToString(M.I.Get("Temp.DirViews", defaultHTMLDir + "viewer/"))

	arrTPL = make([]string, 0)
	arrTPL = append(TplLayout, TplViews + skeleton + "/" + view)

	for _, shared := range shareds {
		arrTPL = append(arrTPL, shared)
	}

	return
}
