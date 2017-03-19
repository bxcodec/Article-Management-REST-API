
# Docs 

[Articles Management API](#) 
Developed under Go (Golang) platform.
This project developed using [gin](https://gopkg.in/gin-gonic/gin.v1) framework to ease the development.
See [Source](#source) For all the library used

## Index

* [Overview](#overview)
* [Support](#support)
* [Quick Start](#quick-start)
* [Source](#source)


## Overview

1. Import the `query.sql`  to your MySql Database. 
2. Set your Database Connections (`username/password` in `services/mysql` and `core/sqlboiler.toml` )
3. Run the project
    
See [Quick Start](#quick-start) for example implementations.


## Support


You can also email <iman.tumorang@gmail.com> or file an [Issue](https://github.com/bxcodec/Article-Management-REST-API/issues/new).


## Quick Start

### Command Line

Run your project using this command

```bash
# download the project into $GOPATH/src
git clone https://github.com/bxcodec/Article-Management-REST-API.git
cd path_of_the_project

# import the query.sql to your mysql server
mysql -u username -p db_name < query.sql

# set your mysql connection authentication
cd "services/mysql"
nano mysql.go

cd "core"
nano sqlboiler.toml

# run the project
go run main.go

```


> If any error happen when running the program caused by missing package, just import it using `go get package_name` command


### Running Your API
```bash
protocol: http
host 	: localhost
port 	: 9090
endpoint: "/v1/articles"

```
For  the API documentation open <http://localhost:9090/v1/docs>

## Source

Below listed all the library used here.

### [gin](https://gopkg.in/gin-gonic/gin.v1) 
Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance, up to 40 times faster thanks to [httprouter](https://github.com/julienschmidt/httprouter). If you need performance and good productivity, you will love Gin.

### [sqlboiler](https://github.com/vattle/sqlboiler)
SQLBoiler is a tool to generate a Go ORM tailored to your database schema.

It is a "database-first" ORM as opposed to "code-first" (like gorm/gorp). That means you must first create your database schema. Please use something like [goose](https://bitbucket.org/liamstask/goose), [sql-migrate](https://github.com/rubenv/sql-migrate) or some other migration tool to manage this part of the database's life-cycle.

### [apidoc](http://apidocjs.com/)
Inline Documentation for RESTful web APIs
apiDoc creates a documentation from API annotations in your source code.