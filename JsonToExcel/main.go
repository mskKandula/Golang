package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tealeg/xlsx/v3"
	"github.com/tidwall/gjson"
)

func main() {

	// studentList := []byte(`[{
	// 	"CandidateDetails":{
	// 		"paperName": "English"
	// 		"programName": "BE"
	// 		"projectTitle": "A Project in English"
	// 		"status": "APPROVED"
	// 		}
	// 		"name": "ABC"
	// 		"userName": "12345678"
	// 	},
	// 	{"CandidateDetails":{
	// 		"paperName": "Science"
	// 		"programName": "B.E"
	// 		"projectTitle": "A Project in Science"
	// 		"status": "REJECTED"
	// 		}
	// 		"name": "DEF"
	// 		"userName": "87654321"
	// 	}
	// 		]`)

	studentList, err := getStudentList()

	if err != nil {
		log.Fatal(err)
	}

	studentListResult := gjson.ParseBytes(studentList)

	Report, err := generateExcel(studentListResult)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your %s report has been generated", Report)

}

func getStudentList() ([]byte, error) {
	res, err := ioutil.ReadFile("jsonFile.json")
	return res, err
}

func generateExcel(studentListResult gjson.Result) (string, error) {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error

	SheetName := "All Students Report"

	file = xlsx.NewFile()

	sheet, err = file.AddSheet(SheetName)

	if err != nil {
		log.Fatal(err)
	}

	row = sheet.AddRow()

	row.SetHeight(15)

	row.Hidden = false

	row.AddCell().SetString("Student Name")

	row.AddCell().SetString("Program Name")

	row.AddCell().SetString("Paper Name")

	row.AddCell().SetString("Project Title")

	row.AddCell().SetString("User Name")

	row.AddCell().SetString("Status")

	ReportName := "Students_Report" + ".xlsx"

	for _, studentDetails := range studentListResult.Array() {

		row := sheet.AddRow()

		row.SetHeight(15)

		row.Hidden = false

		if studentDetails.Get("name").String() == "" {
			row.AddCell().SetString("-")
		} else {

			row.AddCell().SetString(studentDetails.Get("name").String())

		}

		if studentDetails.Get("CandidateDetails").Get("programName").String() == "" {
			row.AddCell().SetString("-")
		} else {

			row.AddCell().SetString(studentDetails.Get("CandidateDetails").Get("programName").String())

		}

		if studentDetails.Get("CandidateDetails").Get("paperName").String() == "" {
			row.AddCell().SetString("-")
		} else {

			row.AddCell().SetString(studentDetails.Get("CandidateDetails").Get("paperName").String())
		}

		if studentDetails.Get("CandidateDetails").Get("projectTitle").String() == "" {
			row.AddCell().SetString("-")
		} else {

			row.AddCell().SetString(studentDetails.Get("CandidateDetails").Get("projectTitle").String())
		}

		if studentDetails.Get("userName").String() == "" {
			row.AddCell().SetString("-")
		} else {

			row.AddCell().SetString(studentDetails.Get("userName").String())
		}

		if studentDetails.Get("CandidateDetails").Get("status").String() == "APPROVED" {

			row.AddCell().SetString("Approved")

		} else if studentDetails.Get("CandidateDetails").Get("status").String() == "REJECTED" {

			row.AddCell().SetString("Rejected")

		} else {

			row.AddCell().SetString("Not Submitted")
		}

	}

	err = file.Save(ReportName)

	if err != nil {
		log.Fatal(err)
	}

	return ReportName, err
}
