package router

import (
	"Blog/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
api.RegisterRouter(r)
}