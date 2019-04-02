package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/qqqasdwx/blog/http/render"
)

func RecoveryHandler(c *gin.Context, err interface{}) {
	render.Data(c, "err", err)
	render.HTML(c, "error")
}
