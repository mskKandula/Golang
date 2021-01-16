package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/tealeg/xlsx/v3"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// defining the excel file headers
var requiredKeys = []string{
	"FirstName",
	"LastName",
	"Address",
	"City",
	"State",
	"Phone No",
	"Email",
	"Web",
}

func main() {

	// reading the excel File
	fileBytes, err := ioutil.ReadFile("SampleFile.xlsx") //Excel file downloaded from internet
	if err != nil {
		log.Fatal(err)
	}

	// passing fileBytes([]bytes) to excel function
	result := excel(fileBytes)

	// just printing the result,It can be either stored in a database / write to a file
	fmt.Println(result)

}

func excel(fileBytes []byte) []gjson.Result {

	// used openBinary method since []bytes
	xlFile, err := xlsx.OpenBinary(fileBytes)
	if err != nil {
		log.Fatal(err)
	}

	// defining data variable as an array of gjson.Result since we return an array of objects
	var objectArray []gjson.Result

	// iterating over sheets in excel file
	for _, sheet := range xlFile.Sheets {

		//  if it has no data
		if sheet.MaxRow < 2 {
			continue
		}

		// iterating over the sheet
		for rowIndex := 1; rowIndex < sheet.MaxRow; rowIndex++ {

			row, _ := sheet.Row(rowIndex)

			values := []interface{}{}

			for i := 0; i < len(requiredKeys); i++ {

				// getting value from a cell
				values = append(values, strings.TrimSpace(row.GetCell(i).String()))

			}

			// calling prepareResult with keys(headers) & values
			object := prepareResult(requiredKeys, values)

			// appending the object to an array
			objectArray = append(objectArray, object)
		}
	}
	return objectArray
}

// maps keys & values to make an object
func prepareResult(keys []string, vals []interface{}) gjson.Result {

	var data string

	for i, k := range keys {

		data, _ = sjson.Set(data, k, vals[i]) //sjson used to set json
	}

	return gjson.Parse(data) //gjson used to get json
}
