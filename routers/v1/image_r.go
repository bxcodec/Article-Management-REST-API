package v1

import (
	"gopkg.in/gin-gonic/gin.v1"

)


func SetImagesRoutes(router *gin.RouterGroup) {
	router.Static("/images", "public/images")
}