package v1

import (

	"gopkg.in/gin-gonic/gin.v1"
)


func InitRoutes(routerGroup *gin.RouterGroup ) {
	
	SetHelloRoutes(routerGroup)
	SetArticlesRoutes(routerGroup)

	SetApiDocRoutes(routerGroup)
	SetImagesRoutes(routerGroup)

}