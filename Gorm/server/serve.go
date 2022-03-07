package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tealeg/xlsx/v3"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var (
	db   *gorm.DB
	err  error
	size int64
)

type User struct {
	// gorm.Model
	Id       uint8 `gorm :"primaryKey;autoIncrement"`
	Name     string
	Email    string
	Mobile   string `gorm :"primaryKey"`
	Password string
}

var requiredKeys = []string{
	"Name",
	"Email",
	"Mobile",
	"Password",
}

// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "please route to /fileUpload to upload file")

// }

func fileUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)

	fmt.Printf("File Size: %+v\n", handler.Size)

	fmt.Printf("MIME Header: %+v\n", handler.Header)

	if strings.Split(handler.Filename, ".")[1] != "xlsx" {
		fmt.Println("File Format Not supported")
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	res := excel(fileBytes)

	for _, val := range res {

		name := val.Get("Name").String()
		email := val.Get("Email").String()
		mobile := val.Get("Mobile").String()
		password := val.Get("Password").String()

		db.Create(&User{Name: name, Email: email, Mobile: mobile, Password: password})
	}
	json.NewEncoder(w).Encode(res)
}

func excel(fileBytes []byte) []gjson.Result {
	xlFile, err := xlsx.OpenBinary(fileBytes)
	if err != nil {
		fmt.Println("error while opening file")
	}
	var data []gjson.Result
	for _, sheet := range xlFile.Sheets {
		if sheet.MaxRow < 2 {
			continue
		}

		for rowIndex := 1; rowIndex < sheet.MaxRow; rowIndex++ {

			row, _ := sheet.Row(rowIndex)

			allKeys := []string{}

			for _, v := range requiredKeys {
				allKeys = append(allKeys, v)
			}
			values := []interface{}{}

			for i := 0; i < len(allKeys); i++ {

				values = append(values, strings.TrimSpace(row.GetCell(i).String()))

			}
			arr := prepareResult(allKeys, values)
			data = append(data, arr)
		}
	}
	return data
}

func prepareResult(keys []string, vals []interface{}) gjson.Result {
	var data string
	for i, k := range keys {
		data, _ = sjson.Set(data, k, vals[i])
	}

	return gjson.Parse(data)
}

func init() {
	db, err = gorm.Open("mysql", "UserName:Password@tcp(127.0.0.1:3306)/sample?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed to Open")
	}

	db.SingularTable(true)

	if db.HasTable(&User{}) == false {
		db.CreateTable(&User{})
	}
}

func main() {

	// db.AutoMigrate(&User{})
	defer db.Close()

	fs := http.FileServer(http.Dir("../client/dist"))
	http.Handle("/", fs)

	http.HandleFunc("/fileUpload", fileUpload)

	fmt.Println("listening on port 8080	")

	http.ListenAndServe(":8080", nil)
}
