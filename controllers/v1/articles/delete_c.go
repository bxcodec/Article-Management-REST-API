package articles

import (
	"gopkg.in/gin-gonic/gin.v1"
	"articles/core/models"
	"articles/services/mysql"
	// "fmt"	
	"strconv"
	"os"
	"strings"
)


func Delete(c *gin.Context) {
	
	var statuscode 		int
	var message			string 		= "Succesfully Deleted"
	var data 			interface{} = nil
	var listError []	models.ErrorModel = nil
	var endpoint		string = c.Request.URL.String()
	var method			string = c.Request.Method

	
	idArticle,_ := strconv.Atoi(c.Param("id"))
	
	article,_:=models.FindArticle(mysql.DBCon, idArticle);
	var err error

	thumbName := article.Thumbnail
	if idArticle != 0{
		err= article.Delete(mysql.DBCon);	
	} 
	if err == nil{

		 dir,_ := os.Getwd()
		path 	:= strings.Replace(dir, " ", "\\ ", -1) + "/public/images/"
		os.Remove(path+thumbName)		
	    
	}

	
	if  err 		!= nil 	||
		article 	== nil	||
		idArticle	== 0{

		var clientErr = false;

		if 	article 	== nil ||
			idArticle 	== 0	{

			var tempEr models.ErrorModel

			tempEr.ErrorCode 	= 4001	
			tempEr.Hints 		= "Article Maybe Already Deleted"
			tempEr.Info			= "visit http://localhost:9090/v1/docs"
			listError 			= append(listError,tempEr)
			statuscode 			= 400	
			clientErr = true
		}		
		if 	err != nil &&
			!clientErr	{
			var tempEr models.ErrorModel
			tempEr.ErrorCode 	= 5001	
			tempEr.Hints 		= "Cannot Delete"
			tempEr.Info			= "visit http://localhost:9090/v1/docs"
			listError 			= append(listError,tempEr)
			statuscode 			= 500	
		}

		message  = "Can't delete the Article"

	}else{
		statuscode = 200
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
