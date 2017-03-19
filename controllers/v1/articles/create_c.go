package articles

import (
	"gopkg.in/gin-gonic/gin.v1"
	"articles/core/models"
	"articles/services/mysql"
	// "fmt"
	// "strconv"
	// "log"
	"io"
 	"strings"
	"os"
 
)






func Create(c *gin.Context) {
	
	
	var statuscode 	int
	var message		string 		= "Successfully Created"
	var data 		interface{} = nil
	var listError [] models.ErrorModel = nil
	var endpoint	string = c.Request.URL.String()
	var method		string = c.Request.Method

	var serverError [3] error

	var request models.RequestArticleModel

	request.Title 	= c.PostForm("title")
	request.Content = c.PostForm("content")


	valid:= validateArticleRequest(request)	


	var article models.Article
	
	article.Title = request.Title
	article.Content = request.Content

	file, header , err1 := c.Request.FormFile("image")


    dir, _ := os.Getwd()
    

	if err1 == nil{

		filename := header.Filename
		article.Thumbnail =filename
	    
	    path := strings.Replace(dir, " ", "\\ ", -1) + "/public/images/"
	    out, err := os.Create(path+filename)
	    
	    serverError[0]= err
	    

	    defer out.Close()

	    _, err = io.Copy(out, file)
	    serverError[1]=err


	}
 	if 	serverError[0] 	==nil	&&
		serverError[1]	==nil	&&
		valid{

			serverError[2] = article.Insert(mysql.DBCon)		
		}
		

	if 	serverError[0] 	!=nil	||
		serverError[1]	!=nil	||
		serverError[2]	!=nil	||
		!valid				 {

			var errClient = false
		
			   if  !valid {
				errClient = true
				var tempEr models.ErrorModel
				tempEr.ErrorCode 	= 4001	
				tempEr.Hints 		= "Required Parameter is Invalid"
				tempEr.Info			= "visit http://localhost:9090/v1/docs"
				listError 			= append(listError,tempEr)
				statuscode 			= 400	
			}

			if 	(serverError[0]		!=nil	||
				 serverError[1]		!=nil	||
				 serverError[2]		!=nil)  && 
				!errClient{

				var tempEr models.ErrorModel
				tempEr.ErrorCode 	= 5001	
				tempEr.Hints 		= "Cannot Create The Article"
				tempEr.Info			= "visit http://localhost:9090/v1/docs"
				listError 			= append(listError,tempEr)
				statuscode 			= 500	
			}
			message= "Cannot Create Article"
	}else{
		statuscode 		=  201
		message			= "Article Created Succesfully"

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

