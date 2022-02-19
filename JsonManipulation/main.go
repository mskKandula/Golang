package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var result string = `[]`

func main() {

	data := `
	[
		{
			"RollNumber" : "2",
			"Name" : "abc",
			"StandardName" : "Semester 1",
			"Submission Date" : "11-02-2022",
			"Obtained Marks" : 13,
			"Total Marks" : 15,
			"Remark" : "Very Good"
		},
		{
			"RollNumber" : "21",
			"Name" : "cde",
			"StandardName" : "Semester 1",
			"Submission Date" : "12-02-2022",
			"Obtained Marks" : 0,
			"Total Marks" : 15,
			"Remark" : ""
		},
		{
			"RollNumber" : "41",
			"Name" : "efg",
			"StandardName" : "Semester 1",
			"Submission Date" : "Not Submitted",
			"Obtained Marks" : 0,
			"Total Marks" : 15,
			"Remark" : ""
		},
		{
			"RollNumber" : "21",
			"Name" : "cde",
			"StandardName" : "Semester 1",
			"Submission Date" : "12-02-2022",
			"Obtained Marks" : 0,
			"Total Marks" : 15,
			"Remark" : ""
		}
	]`

	file, err := os.Create("UpdatedData.txt")
	if err != nil {
		log.Fatal(err)
	}

	jsonArray := gjson.Parse(data).Array()

	for _, jsonObj := range jsonArray {

		submissionDate := jsonObj.Get("Submission Date").String()
		obtainedMarks := jsonObj.Get("Obtained Marks").Int()
		remark := jsonObj.Get("Remark").String()
		stringData := fmt.Sprintf("%v", jsonObj)

		if submissionDate == "Not Submitted" {

			stringData, err = sjson.Set(stringData, "isEvaluated", false)
			if err != nil {
				log.Println(err)
			}
		} else if submissionDate != "Not Submitted" && obtainedMarks == 0 && remark == "" {
			stringData, err = sjson.Set(stringData, "isEvaluated", false)
			if err != nil {
				log.Println(err)
			}
		} else {
			stringData, err = sjson.Set(stringData, "isEvaluated", true)

			if err != nil {
				log.Println(err)
			}
		}

		result, err = sjson.Set(result, "-1", stringData)
		if err != nil {
			log.Fatal(err)
		}
	}

	n, err := file.WriteString(fmt.Sprintf("%v", gjson.Parse(result).Value()))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total Bytes Written :", n)
}
