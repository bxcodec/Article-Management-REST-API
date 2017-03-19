package hello

import (
	"gopkg.in/gin-gonic/gin.v1"
	"articles/core/models"
	
	
)



func HelloWorld(ctx *gin.Context) {
	
	var statuscode 	int
	var message		string 		= "Hello World"
	var data 		interface{} = nil
	var listError [] models.ErrorModel = nil
	var endpoint	string = ctx.Request.URL.String()
	var method		string = ctx.Request.Method
	
    

	responseModel := &models.ResponseModel{
		statuscode,
		message,
		data,
		listError,
		endpoint,
		method,
	} 

	var content gin.H = responseModel.NewResponse(); 
	



	ctx.JSON(200,content)
}