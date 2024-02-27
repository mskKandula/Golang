package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	xlsx "github.com/tealeg/xlsx/v3"
)

type QuizOption struct {
	OptionText   string `json:"optionText"`
	FeedbackText string `json:"feedbackText"`
}

// type Stack []string

// // IsEmpty: check if stack is empty
// func (s *Stack) IsEmpty() bool {
// 	return len(*s) == 0
// }

// // Push a new value onto the stack
// func (s *Stack) Push(str string) {
// 	*s = append(*s, str) // Simply append the new value to the end of the stack
// }

// // Remove and return top element of stack. Return false if stack is empty.
// func (s *Stack) Pop() (string, bool) {
// 	if s.IsEmpty() {
// 		return "", false
// 	} else {
// 		index := len(*s) - 1   // Get the index of the top most element.
// 		element := (*s)[index] // Index into the slice and obtain the element.
// 		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
// 		return element, true
// 	}
// }

// func modifyString(src string) string {

// 	var stack Stack
// 	// Input JSON string
// 	newString := ""
// 	for _, s := range src {

// 		if string(s) == `"` {
// 			if stack.IsEmpty() {
// 				stack.Push(`"`)
// 			} else {
// 				stack.Pop()
// 			}
// 		}

// 		if (string(s) == "\n" || string(s) == "\t") && stack.IsEmpty() {
// 			continue

// 		}

// 		newString += string(s)

// 	}

// 	return newString

// }

func modifyString(src string) string {
	var insideQuotes bool
	var modifiedString string

	for _, char := range src {
		if char == '"' {
			insideQuotes = !insideQuotes
		}

		if char == '\n' && !insideQuotes {
			continue
		}

		modifiedString += string(char)
	}

	return modifiedString
}

func main() {
	// Input JSON string
	// jsonString := "[{\n    \"optionText\": \"class Student \n    { \n      int[] scores = new int[5] {3, 2, 4, 1, 5};  \n      public int this[ int index ]\n      {\n        set {  \n          if (index \u003c 5) scores[index] = value;  \n          else Console.WriteLine(\\\"Invalid\\n Index\\\"); \n          }  \n      }  \n    }\",\n    \"feedbackText\": \"\"\n  },\n  {\n    \"optionText\": \"class Student \n    { \n      int[] scores = new int[5] {3, 2, 4, 1, 5\n      }; \n      public int this[ int index ] \n      {  \n        get \n        {  \n          if (index \u003c 5) return scores[ index ];  \n          else \n          {  \n            Console.WriteLine(\\\"Invalid\\n Index\\\"); return 0;\n            }  \n        }  \n        set \n        {  \n\t  if (index \u003c 5) scores[ index ] = value; else  Console.WriteLine(\\\"Invalid Index\\\");  \n        }  \n      }  \n    }\",\n    \"feedbackText\": \"\"\n  },\n  {\n    \"optionText\": \"class Student { int[] scores = new int[5] {3, 2, 4, 1, 5};  public int this[ int index ] {  get {  if (index \u003c 5) return scores[ index ];  else {  Console.WriteLine(\\\"Invalid Index\\\");  return 0;  }  }  }  }\",\n    \"feedbackText\": \"\"\n  },\n  {\n    \"optionText\": \"class Student { int[] scores = new int[5] {3, 2, 4, 1, 5};  public int this[ int index ] {  get { if (index \u003c 5) scores[ index ] = value;  else {  Console.WriteLine(\\\"Invalid Index\\\"); }  } set {  if (index \u003c 5) return scores[ index ]; else {  Console.WriteLine(\\\"Invalid Index\\\"); return 0; } } } }\",\n    \"feedbackText\": \"\"\n  }]"
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

	result := modifyString(jsonString)

	// Replace unescaped newline characters with escaped newline characters
	result = strings.ReplaceAll(result, "\n", "\\n")
	result = strings.ReplaceAll(result, "\t", "\\t")

	var quizOptions []QuizOption

	// Unmarshal the JSON string into the struct
	err = json.Unmarshal([]byte(result), &quizOptions)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
	}

	// Display the decoded quiz options
	fmt.Println("Decoded Quiz Options:")
	for _, decodedOption := range quizOptions {
		fmt.Printf("Option Text: %s\nFeedback Text: %s\n\n", decodedOption.OptionText, decodedOption.FeedbackText)
	}

}
