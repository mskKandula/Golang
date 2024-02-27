package main

import (
	"encoding/json"
	"fmt"
	"log"

	xlsx "github.com/tealeg/xlsx/v3"
)

type QuizOption struct {
	OptionText   string `json:"optionText"`
	FeedbackText string `json:"feedbackText"`
}

func main() {
	excelFileName := "quiz_data.xlsx" // Update this with your Excel file name

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Error opening Excel file: %s", err)
	}

	// Assuming the JSON data is in the first cell (A1) of the first sheet (index 0)
	sheet := xlFile.Sheets[0]
	jsonCell, _ := sheet.Cell(0, 0) // Cell at row 1, column 1 (0-indexed)

	// Get the JSON string from the cell
	jsonString := jsonCell.String()

	// jsonBytes, _ := json.Marshal(jsonString)
	// fmt.Println(string(jsonBytes))
	// jsonString = strings.ReplaceAll(jsonString, ",\n", ",")
	// jsonString = strings.ReplaceAll(jsonString, "{\n", "{")
	// jsonString = strings.ReplaceAll(jsonString, `""\n`, `""`)

	// jsonString = strings.ReplaceAll(jsonString, "\t", "")

	// fmt.Println("Raw String :")
	// fmt.Println(jsonString)

	jsonBytes, _ := json.Marshal(jsonString)

	fmt.Println("Marshal bytes in string:")
	fmt.Println(string(jsonBytes))

	var quizOptions []QuizOption

	// Unmarshal the JSON string into the struct
	err = json.Unmarshal([]byte(jsonString), &quizOptions)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
	}

	// Display the decoded quiz options
	fmt.Println("Decoded Quiz Options:")
	for _, decodedOption := range quizOptions {
		fmt.Printf("Option Text: %s\nFeedback Text: %s\n\n", decodedOption.OptionText, decodedOption.FeedbackText)
	}
}
