package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	R "github.com/yuw-mvc/yuw/routes"
	"net/http"
)

func Redirect(ctx *gin.Context, location string) {
	ctx.Abort()

	d, ok := R.RMaP[location]
	if ok {
		requestURL := "http://" + ctx.Request.Host + cast.ToString(d)
		ctx.Redirect(http.StatusFound, requestURL)
	}
}
