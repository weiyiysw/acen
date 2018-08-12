package main

import (
	"acen/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	router.Router(&app.RouterGroup)

	app.Run(":9825")
}
