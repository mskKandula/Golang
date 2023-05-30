package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mskKandula/SPH/controller"
)

var err error

func init() {
	// "root:root@tcp(db:3306)/SPH"
	controller.Db, err = sql.Open("mysql", "root:connect@123@tcp(localhost:3306)/SPH")

	if err != nil {
		log.Fatal("DB Connection Failed to Open")
	}
}

func main() {

	r := gin.Default()
	r.POST("/articles", controller.CreateArticle)
	r.GET("/articles/:id", controller.GetArticleById)
	r.GET("/articles", controller.GetAllArticles)
	r.Run(":8080")
}
