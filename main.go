package main

import (
	"Blog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
    r.LoadHTMLGlob("templates/*")
	r.Static("/assets","./assets")
	router.InitRouter(r)
	r.Run(":8089")
}