package cors

import (
        "gopkg.in/gin-gonic/gin.v1"

)


func CORSMiddleware() gin.HandlerFunc {

     return func(c *gin.Context) {

         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
         c.Writer.Header().Set("Access-Control-Max-Age", "86400")
         c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, DELETE, UPDATE")
         c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
         c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

         if c.Request.Method == "OPTIONS" {
             c.AbortWithStatus(204)
         } else {
             c.Next()
         }
     }
 }

 
// func CORSMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     }
// }





