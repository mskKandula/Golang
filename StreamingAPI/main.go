package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	bufChan = make(chan string, 10)
)

func main() {
	go asyncHLSConversion(bufChan)

	defer func() {
		close(bufChan)
	}()
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

	bufChan <- path

	c.JSON(http.StatusOK, gin.H{"fileUploaded": "Success"})

}

// file path creation
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func asyncHLSConversion(filePaths <-chan string) {

	for filePath := range filePaths {
		paths := strings.Split(filePath, "/")

		path := paths[0] + "/" + paths[1]

		err := os.Chdir(path)

		if err != nil {
			fmt.Println(err.Error())
		}

		cmd := exec.Command("ffmpeg", "-i", paths[2], "-codec:", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", "index.m3u8")

		err = cmd.Run()

		if err != nil {
			fmt.Println(err.Error())
		}

		imageFileName := paths[2] + ".png"

		cmd = exec.Command("ffmpeg", "-i", paths[2], "-ss", "00:00:01.000", "-vframes", "1", imageFileName)

		err = cmd.Run()

		if err != nil {
			fmt.Println(err.Error())
		}

		err = os.Remove(paths[2])

		if err != nil {
			fmt.Println(err.Error())
		}

		err = os.Chdir("../..")
		if err != nil {
			fmt.Println(err.Error())
		}

	}
}
