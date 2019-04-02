package test

import (
	"github.com/gin-gonic/gin"
)

func TestPanic(c *gin.Context) {
	panic("test panic")
}
