package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/handle", fileHandler)
	r.Run(":8081")
}

func fileHandler(c *gin.Context) {
	file, handler, err := c.Request.FormFile("myFile")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer file.Close()

	paths := strings.Split(handler.Filename, ".")

	if paths[1] != "mp4" {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported File Format"})
		return
	}

	if handler.Size > 10*1024*1024 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "File size is big"})
		return
	}

	path := "media/" + paths[0] + "/" + handler.Filename

	dstFile, err := create(path)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = io.Copy(dstFile, file)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer dstFile.Close()

	c.JSON(http.StatusOK, gin.H{"fileUploaded": "Success"})

}
