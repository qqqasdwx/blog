package unknown

import (
	"github.com/gin-gonic/gin"
	"github.com/qqqasdwx/blog/http/render"
)

func NoRoute(c *gin.Context) {
	render.HTML(c, "404")
}

func NoMethod(c *gin.Context) {
	render.HTML(c, "404")
}
