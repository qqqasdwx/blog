package http

import (
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/qqqasdwx/blog/config"
	"github.com/qqqasdwx/blog/http/middleware"
	"github.com/qqqasdwx/blog/http/render"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func Run() {
	render.Init()

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(nice.Recovery(middleware.RecoveryHandler))

	router.Static("/static", "static/")
	ConfigRoute(router)

	router.Run(config.Config().Http.Listen)
}
