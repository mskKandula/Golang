package main

//added
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/sjson"
)

const (
	Url string = "http://1232.com"
)

func main() {

	fmt.Println("Hello")

	LearnerCodeList := [][]string{{"123", "546", "789"}, {"546", "342", "265"}}

	CenterCodeList := []string{"4352r435refr", "654t53482"}

	for index, CenterCode := range CenterCodeList {

		requestObj := `{}`

		requestObj, _ = sjson.Set(requestObj, CenterCode, LearnerCodeList[index])

		req, err := http.NewRequest("POST", Url, bytes.NewBuffer([]byte(requestObj)))

		if err != nil {
			log.Fatal(err.Error())
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("response Body:", string(body))

		//Prepare the data

		//Service call to store in DB
	}
}
