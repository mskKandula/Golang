package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", stringHandler)
	fmt.Println("listening on port 8081")
	r.Run(":8082")

}

func stringHandler(c *gin.Context) {
	if true {
		log.Fatal("error")
	}
	c.JSON(200, gin.H{"obj": "another one"})

}
