package v1

import (
	"gopkg.in/gin-gonic/gin.v1"
	"articles/controllers/v1/hello"
)


func SetHelloRoutes(router *gin.RouterGroup) {
	
	router.GET("/hello",hello.HelloWorld)
}