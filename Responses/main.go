package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var (
	imagePathChannel        = make(chan string, 500)
	rootPath         string = "../ProcessingPipeline/ImageProcessing/Images/"
	// files            []string
)

func main() {
	http.HandleFunc("/", stringHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/jsonVariant", jsonVariantHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/zipFile", zipFileHandler)
	fmt.Println("listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

}

//stringHandler returns http respone in string format.
func stringHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Buddy!")

}

//jsonHandler returns http respone in JSON format.
func jsonHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var u User

	// Unmarshalling the data into User struct
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// do some manipulations on User struct

	// sending the final User struct as a response
	json.NewEncoder(w).Encode(u)

}

func jsonVariantHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var name, last string

	response, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonList := gjson.ParseBytes(response)

	var dataToSend []gjson.Result

	for _, val := range jsonList.Array() {
		name = val.Get("name").String()
		last = val.Get("last").String()
		data := prepareResult([]string{"name", "last"}, []interface{}{name, last})
		dataToSend = append(dataToSend, data)
	}
	json.NewEncoder(w).Encode(dataToSend)
}

func prepareResult(keys []string, vals []interface{}) gjson.Result {
	var data string
	for i, k := range keys {
		data, _ = sjson.Set(data, k, vals[i])
	}

	return gjson.Parse(data)
}

// fileHandler downloads a file on the client machine
func fileHandler(w http.ResponseWriter, r *http.Request) {

	imageFile, err := os.Open("./Sample.png")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// its for the file to download automatically as "Sample.png"
	w.Header().Set("Content-Disposition", "attachment; filename=Sample.png")

	w.Header().Set("Content-Type", "image/png")

	defer imageFile.Close()

	// copying the imageFile to response writer
	io.Copy(w, imageFile)
}

//templateHandler renders a template and returns as http response.
func templateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := User{Id: 1,
		Name:  "YourName",
		Email: "YourEmail",
		Phone: "YourNumber"}
	t.Execute(w, user)
}

func zipFileHandler(w http.ResponseWriter, r *http.Request) {

	filename := "ImageFiles"

	buf := new(bytes.Buffer)

	zipWriter := zip.NewWriter(buf)

	go func() {
		defer close(imagePathChannel)
		if err := filepath.Walk(rootPath, visit(imagePathChannel)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}()

	for file := range imagePathChannel {

		fullPath := strings.Split(file, string("\\"))

		specificPath := strings.Join(fullPath[1:], "/")

		pathWriter, err := zipWriter.Create(specificPath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		byteData, err := os.ReadFile(file)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = pathWriter.Write(byteData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	zipWriter.Close()
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", filename))
	w.Write(buf.Bytes())
}

func visit(imagePathChannel chan<- string) filepath.WalkFunc {

	return func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatal(err)
		}

		// specifying the paths of files which are of png format
		if filepath.Ext(path) == ".png" {

			imagePathChannel <- path

		}
		return nil
	}
}
