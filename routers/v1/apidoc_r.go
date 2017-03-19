package v1

import (
	"gopkg.in/gin-gonic/gin.v1"

)


func SetApiDocRoutes(router *gin.RouterGroup) {
	router.Static("/docs", "public/apidoc")
}