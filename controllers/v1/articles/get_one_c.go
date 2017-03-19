package articles

import (
	"gopkg.in/gin-gonic/gin.v1"
	"articles/core/models"
	"articles/services/mysql"
	// "fmt"	
	"strconv"
)


type ResponseArticle struct{
	ThumbUrl string `json:"thumb_url"`
	Article  *models.Article `json:"article"`
}

func GetOne(c *gin.Context) {
	
	var statuscode 		int
	var message			string 		= "Hello World"
	var data 			interface{} = nil
	var listError []	models.ErrorModel = nil
	var endpoint		string = c.Request.URL.String()
	var method			string = c.Request.Method

	
	idArticle,_ := strconv.Atoi(c.Param("id"))
	
	article,err:=models.FindArticle(mysql.DBCon, idArticle);
	

	if  err 		!= nil 	||
		article 	== nil	||
		idArticle	== 0{

		var clientErr = false;

		if 	article 	== nil ||
			idArticle 	== 0	{

			var tempEr models.ErrorModel
			tempEr.ErrorCode 	= 4001	
			tempEr.Hints 		= "Article Not Found Article Maybe Already Deleted"
			tempEr.Info			= "visit http://localhost:9090/v1/docs"
			listError 			= append(listError,tempEr)
			statuscode 			= 400	
			clientErr = true
		}		
		if 	err != nil &&
			!clientErr	{
			var tempEr models.ErrorModel
			tempEr.ErrorCode 	= 5001	
			tempEr.Hints 		= "Cannot Find The Article"
			tempEr.Info			= "visit http://localhost:9090/v1/docs/"
			listError 			= append(listError,tempEr)
			statuscode 			= 500	
		}

		message  = "Can not Find The Article "

	}else{

		var res ResponseArticle
		res.ThumbUrl = "http://"+c.Request.Host+"/v1/images/"
		res.Article  = article 
		statuscode = 200
		data = res
		message  = "Succesfully Retrieved"		
	}

	responseModel := &models.ResponseModel{
		statuscode,
		message,
		data,
		listError,
		endpoint,
		method,
	} 

	var content gin.H = responseModel.NewResponse(); 
	

	c.JSON(statuscode,content)
}
