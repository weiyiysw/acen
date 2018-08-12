package router

import (
	"acen/handlers/hello"

	"github.com/gin-gonic/gin"
)

// Router define
func Router(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		v1.GET("/hello", hello.Hello)
	}
}
