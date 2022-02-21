package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

var rows []string

func main() {

	studentList := `[{
			"paperName": "English",
			"programName": "BE",
			"projectTitle": "A Project in English",
			"status": "APPROVED",
			"name": "ABC",
			"userName": "12345678"
		},
		{
			"paperName": "English",
			"programName": "BE",
			"projectTitle": "A Project in English",
			"status": "APPROVED",
			"name": "ABC",
			"userName": "12345678"
		}
		]`

	SheetName := "All Students Report"

	ReportName := "Students_Report" + ".xlsx"

	var result []map[string]interface{}

	json.Unmarshal([]byte(studentList), &result)

	Report, err := generateExcel(result, SheetName, ReportName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your %s report has been generated", Report)
}

func generateExcel(studentListResult []map[string]interface{}, SheetName, ReportName string) (string, error) {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error

	file = xlsx.NewFile()

	sheet, err = file.AddSheet(SheetName)

	if err != nil {
		log.Fatal(err)
	}

	row = sheet.AddRow()

	row.SetHeight(15)

	row.Hidden = false

	for key := range studentListResult[0] {

		row.AddCell().SetString(key)
		rows = append(rows, key)

	}

	for _, obj := range studentListResult {

		row = sheet.AddRow()

		row.SetHeight(15)

		row.Hidden = false

		for _, key := range rows {

			val := obj[key]

			row.AddCell().SetString(val.(string))

		}

	}

	err = file.Save(ReportName)

	if err != nil {
		log.Fatal(err)
	}

	return ReportName, err
}
