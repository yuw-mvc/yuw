package controllers

import (
	"github.com/gin-gonic/gin"
	E "github.com/yuw-mvc/yuw/exceptions"
	"github.com/yuw-mvc/yuw/mvc/middleware"
	"github.com/yuw-mvc/yuw/mvc/services"
	"net/http"
)

type (
	Controllers struct {
		Srv *services.Services
	}
)

func New() *Controllers {
	return &Controllers {
		Srv: services.New(),
	}
}

func (c *Controllers) Redirect(ctx *gin.Context, location string) {
	middleware.Redirect(ctx, location)
}

func (c *Controllers) To(ctx *gin.Context, res *services.PoT) {
	switch res.Type {
	case services.ToJSON:
		ctx.JSON(
			res.Code,
			res.Response,
		)

		return

	case services.ToHTML:
		if res.HTML == "" {
			ctx.AbortWithError(
				http.StatusNotFound,
				E.Err("yuw^error",E.ErrPosition()),
			)
		} else {
			ctx.HTML(
				res.Code,
				res.HTML,
				res.Response,
			)
		}

		return

	default:
		ctx.AbortWithError(
			http.StatusNotFound,
			E.Err("yuw^error",E.ErrPosition()),
		)

		return
	}
}

