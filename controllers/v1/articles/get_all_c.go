package articles

import (
	"gopkg.in/gin-gonic/gin.v1"
	"articles/core/models"
	"articles/services/mysql"
	. "github.com/vattle/sqlboiler/queries/qm"
	"fmt"	
	"strconv"
)

type ResponseManyArticle struct{
	ThumbUrl string `json:"thumb_url"`
	TotalArticle int  `json:"total"`
	Article  []*models.Article `json:"articles"`
}


type QueryArticles struct{
	Title		string
	Page		int
}

func GetAll(c *gin.Context) {
	
	var statuscode 	int
	var message		string 		= "Get The Articles"
	var data 		interface{} = nil
	var listError [] models.ErrorModel = nil
	var endpoint	string = c.Request.URL.String()
	var method		string = c.Request.Method

	var query QueryArticles
	
	query.Title						=	"%"+ c.Query("q") + "%"
	query.Page,_					= 	strconv.Atoi(c.Query("p"))

	var err error 
	var articles []*models.Article

	if query.Page > 0{
		query.Page -=1
	}
	
	articles,err=models.Articles(mysql.DBCon, 
						 Where("title like ?",query.Title),
						 Limit(10),
						 Offset(query.Page*10)).All()	
						 fmt.Println(err)	
	
	
	if  err	!= nil 	{	
		var tempEr models.ErrorModel
		tempEr.ErrorCode 	= 5001	
		tempEr.Hints 		= "Cannot Fetch The Article"
		tempEr.Info			= "visit http://localhost:9090/v1/docs"
		listError 			= append(listError,tempEr)
		statuscode 			= 500	
		message  = "Can not Fetch The Articles "

	}else{
		statuscode = 200
		
		var res ResponseManyArticle
		res.ThumbUrl = "http://"+c.Request.Host+"/v1/images/"
		res.Article  = articles
		res.TotalArticle = len(articles)  
		data = res;

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
