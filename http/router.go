package http

import (
	"github.com/gin-gonic/gin"
	testhandler "github.com/qqqasdwx/blog/http/handler/test"
	"github.com/qqqasdwx/blog/http/handler/unknown"
)

func ConfigRoute(router *gin.Engine) {
	configTestRoute(router)
	configUnknownRoute(router)
}

func configTestRoute(router *gin.Engine) {
	test := router.Group("/test")
	test.GET("/panic", testhandler.TestPanic)
}

func configUnknownRoute(router *gin.Engine) {
	router.NoRoute(unknown.NoRoute)
	router.NoMethod(unknown.NoMethod)
}
