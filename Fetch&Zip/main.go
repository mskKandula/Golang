package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type FileData struct {
	Name string
	Data []byte
}

func main() {
	fmt.Println("Please Wait.....")

	var urls = []string{"https://era.mkcl.org/lms/cdnserver/1sZIq76KXCy2J3fFnDnR85OVSpY/IN272500101/Student/gurunathm1/StudentAssignments/03/1 Answers for Cyber Security_T1677233311182.pdf",
		"https://era.mkcl.org/lms/cdnserver/1sZIq76KXCy2J3fFnDnR85OVSpY/IN272500101/Student/apurvav1/StudentAssignments/03/Answers for Cyber Security_T1677233283981.pdf",
		"https://era.mkcl.org/lms/cdnserver/1sZIq76KXCy2J3fFnDnR85OVSpY/IN272500101/Student/apurvav1/StudentAssignments/01/Session- 1 Assignment_answer_T1677233202783.pdf",
		"https://era.mkcl.org/lms/cdnserver/1sZIq76KXCy2J3fFnDnR85OVSpY/IN272500101/Student/gurunathm1/StudentAssignments/01/Session- 1 Assignment_ans_T1677232801371.pdf"}

	bytesChan := make(chan FileData, 3)
	signalChan := make(chan struct{})

	go converter(bytesChan, signalChan)

	fetchFiles(urls, bytesChan)
	<-signalChan
}

func fetchFiles(urls []string, bytesChan chan<- FileData) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string, wg *sync.WaitGroup) {

			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
			}

			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}

			strArr := strings.Split(url, "/")

			fileName := strArr[len(strArr)-1]

			bytesChan <- FileData{
				fileName,
				bodyBytes,
			}

			wg.Done()

		}(url, &wg)

	}
	wg.Wait()
	close(bytesChan)
}

func converter(bytesChan <-chan FileData, signalChan chan<- struct{}) {

	buf := new(bytes.Buffer)

	// Create a new zip archive.
	zipWriter := zip.NewWriter(buf)

	for fileData := range bytesChan {

		zipFile, err := zipWriter.Create(fileData.Name)

		if err != nil {
			log.Println(err)
		}

		_, err = zipFile.Write(fileData.Data)

		if err != nil {
			log.Println(err)
		}
	}

	err := zipWriter.Close()

	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile("sample.zip", buf.Bytes(), 0777)

	if err != nil {
		log.Println(err)
	}
	signalChan <- struct{}{}
}
