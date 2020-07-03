package views

import (
	"github.com/spf13/cast"
	M "github.com/yuw-mvc/yuw/modules"
	"github.com/yuw-mvc/yuw/routes"
)

type Views struct {
	U *M.Utils
}

func New() *Views {
	return &Views {
		U: M.NewUtils(),
	}
}

func (v *Views) Url(i string) (res string) {
	strHost := cast.ToString(M.I.Get("Yuw.Url", ""))

	_, ok := routes.RMaP[i]
	if ok {
		res = strHost + cast.ToString(routes.RMaP[i])
	} else {
		res = strHost + cast.ToString(i)
	}

	return
}

func (v *Views) StaticURL() string {
	strPost := cast.ToString(M.I.Get("Yuw.port", "8888"))
	defaultStaticURL := "http://127.0.0.1:" + strPost

	return cast.ToString(M.I.Get("Yuw.StaticURL", defaultStaticURL))
}


