package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/mskKandula/SPH/model"
)

var (
	Db  *sql.DB
	err error
)

func CreateArticle(c *gin.Context) {

	article := model.Article{}
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error(), "data": nil})
		return
	}

	query, err := Db.Prepare("INSERT INTO Articles(title, content, author) VALUES(?,?,?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error(), "data": nil})
		return
	}

	res, err := query.Exec(article.Title, article.Content, article.Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error(), "data": nil})
		return
	}
	resId, err := res.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Success", "data": gin.H{"id": resId}})
}

func GetArticleById(c *gin.Context) {
	article := model.Article{}

	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error(), "data": nil})
		return
	}

	row := Db.QueryRow("select * from Articles where id=?", intId)

	err = row.Scan(&article.Id, &article.Title, &article.Content, &article.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error(), "data": nil})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success", "data": article})
}

func GetAllArticles(c *gin.Context) {

	articles := []model.Article{}

	rows, err := Db.Query("select * from Articles")

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error(), "data": nil})
			return
		}
	}

	defer rows.Close()

	for rows.Next() {
		var article model.Article

		if err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Author); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error(), "data": nil})
			return

		}
		articles = append(articles, article)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success", "data": articles})
}
