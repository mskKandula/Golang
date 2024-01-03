package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

type PythonProgram struct {
	Code string `json:"code"`
}

func main() {
	excelFileName := "python_programs.xlsx" // Update this with your Excel file name

	// Open the Excel file
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Error opening Excel file: %s", err)
	}

	var programs []PythonProgram

	// Assuming the Python programs are in the first sheet (index 0)
	sheet := xlFile.Sheets[0]

	// Iterate through rows in the sheet
	for _, row := range sheet.Rows {
		// Assuming the Python programs are in the first column (index 0)
		cell := row.Cells[0]
		pythonProgram := cell.String()

		// Create PythonProgram struct
		program := PythonProgram{
			Code: pythonProgram,
		}

		// Append program to programs slice
		programs = append(programs, program)
	}

	// fmt.Println(programs)

	// Marshal the programs into JSON
	jsonData, err := json.Marshal(programs)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %s", err)
	}

	// Display the marshaled JSON
	fmt.Println(string(jsonData))

	// Unmarshal the JSON back to the struct
	var decodedPrograms []PythonProgram
	err = json.Unmarshal(jsonData, &decodedPrograms)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
	}

	// Display the decoded Python programs
	fmt.Println("Decoded Python Programs:")
	for _, decodedProgram := range decodedPrograms {
		fmt.Println(decodedProgram.Code)
	}
}
