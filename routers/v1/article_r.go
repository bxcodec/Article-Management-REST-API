package v1

import (
	"gopkg.in/gin-gonic/gin.v1"
	article "articles/controllers/v1/articles"

)


func SetArticlesRoutes(router *gin.RouterGroup) {
/**
 * @apiDefine ApiErrors
 * @apiError (Error 4xxx) {4001} BadRequest Some required parameter is invalid.
 * @apiError (Error 5xxx) {5001} InternalServerError Something wrong with our services
 */

	



/**
 * @api {get} /v1/articles Get Many
 * @apiDescription
 * Used for get many article (max = 10 in one page) <br>
 * Able to querying with title of the articles
 * @apiGroup Articles
 * @apiParam {String} [q] Input query for title
 * @apiParam {String} [p] Page of the content
 * @apiExample Example usage:
 * curl -XGET http://localhost:9090/v1/articles
 * curl -XGET http://localhost:9090/v1/articles?p=2
 * curl -XGET http://localhost:9090/v1/articles?q=makan
 * curl -XGET http://localhost:9090/v1/articles?q=makan&p=1
 * @apiSuccessExample {json} Success
{
  "data": {
    "thumb_url": "http://localhost:9090/v1/images/",
    "total": 10,
    "articles": [
      {
        "id_articles": 1,
        "title": "Kala Senja Sejuk",
        "content": "Ketika malam tiba ,saya menangis",
        "thumbnail": "thumb.png",
        "created_at": null,
        "updated_at": null
      },
      {
        "id_articles": 2,
        "title": "Hello Adik Adik",
        "content": "Makan kue enak sekali",
        "thumbnail": "",
        "created_at": "2017-03-16T15:30:20Z",
        "updated_at": "2017-03-16T15:30:20Z"
      },
     //... 
    ]
  },
  "endpoint": "/v1/articles",
  "errors": null,
  "message": "Get The Articles",
  "method": "GET",
  "statuscode": 200,
  "timestamp": "2017-03-17 08:13:17"
}
 *@apiUse ApiErrors
 */
	router.GET("/articles",article.GetAll)

/**
 * @api {post} /v1/articles Create Article
 * @apiDescription
 * Create an Article
 * @apiGroup Articles
 * @apiParam {String} title The title of the Article
 * @apiParam {String} content The content of the Article
 * @apiParam {multipart/form-data} image The thumbnail of the Article
 * @apiExample Example usage:
 * curl -XPOST -F "title=Makan Kue Yang Enak"   -F "content=Hello Dunia Indah"  -F "image=@/home/ww3/Pictures/oke.png" http://localhost:9090/v1/articles
 * @apiSuccessExample {json} Success
{
  "data": null,
  "endpoint": "/v1/articles",
  "errors": null,
  "message": "Article Created Succesfully",
  "method": "POST",
  "statuscode": 201,
  "timestamp": "2017-03-17 09:20:10"
}	
 *@apiUse ApiErrors
 */	
	router.POST("/articles",article.Create)
/**
 * @api {get} /v1/articles/:id Get One
 * @apiDescription
 * Get One Article
 * @apiGroup Articles
 * @apiParam {Integer} id Id Of the Article
 * @apiExample Example usage:
 * curl -XGET localhost:9090/v1/articles/1
 * @apiSuccessExample {json} Success
{
  "data": {
    "thumb_url": "http://localhost:9090/v1/images/",
    "article": {
      "id_articles": 1,
      "title": "Kala Senja Sejuk",
      "content": "Ketika malam tiba ,saya menangis",
      "thumbnail": "thumb.png",
      "created_at": null,
      "updated_at": null
    }
  },
  "endpoint": "/v1/articles/1",
  "errors": null,
  "message": "Succesfully Retrieved",
  "method": "GET",
  "statuscode": 200,
  "timestamp": "2017-03-17 09:25:04"
}
 *@apiUse ApiErrors
 */

	router.GET("/articles/:id",article.GetOne)
/**
 * @api {delete} /v1/articles/:id Delete an Article
 * @apiDescription
 * Hard Delete an Article
 * @apiGroup Articles
 * @apiParam {Interger} id Id of the article
 * @apiExample Example usage:
 * curl -XDELETE localhost:9090/v1/articles/5
 * @apiSuccessExample {json} Success
{
  "data": null,
  "endpoint": "/v1/articles/6",
  "errors": null,
  "message": "Succesfully Deleted",
  "method": "DELETE",
  "statuscode": 200,
  "timestamp": "2017-03-17 09:29:21"
}	
 *@apiUse ApiErrors
 */	
	router.DELETE("/articles/:id",article.Delete)

/**
 * @api {put} /v1/articles/:id Update Article
 * @apiDescription
 * Update an Article
 * @apiGroup Articles
 * @apiParam {Integer} id Id Of The Article
 * @apiParam {String} title The title of the Article
 * @apiParam {String} content The content of the Article
 * @apiParam {multipart/form-data} image The thumbnail of the Article
 * @apiExample Example usage:
 * curl -XPUT -F "title=Makan Kue Yang Enak"   -F "content=Hello Dunia Indah"  -F "image=@/home/ww3/Pictures/oke.png" http://localhost:9090/v1/articles/7
 * @apiSuccessExample {json} Success
{
  "data": null,
  "endpoint": "/v1/articles/7",
  "errors": null,
  "message": "Article Updated Succesfully",
  "method": "PUT",
  "statuscode": 200,
  "timestamp": "2017-03-17 09:30:14"
}	
 *@apiUse ApiErrors
 */	
	router.PUT("/articles/:id",article.Update)

}