package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

var DBCon 	*sql.DB
var err 	error


func init() {
 ConnectDB();	
}

func ConnectDB() {
	//CHANGE HERE WITH YOUR MYSQL AUTHENTICATION
	DBCon,err =sql.Open("mysql","root:15pstr@tcp(127.0.0.1:3306)/iman_db?parseTime=true")
}


func CheckDB() *sql.DB  {
	
	if   pinger  := DBCon.Ping() ; err != nil || pinger!=nil  {
  			fmt.Println("FAILDE TO CONNECT", "err" , err)
  			
		}
	
	return DBCon
}