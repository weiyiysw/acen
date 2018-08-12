package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello handler
func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
