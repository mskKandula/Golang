package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type QuizOption struct {
	OptionText   string `json:"optionText"`
	FeedbackText string `json:"feedbackText"`
}

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
	jsonString := "[{\n    \"optionText\": \"class Student \n    { \n      int[] scores = new int[5] {3, 2, 4, 1, 5};  \n      public int this[ int index ]\n      {\n        set {  \n          if (index \u003c 5) scores[index] = value;  \n          else Console.WriteLine(\\\"Invalid\\n Index\\\"); \n          }  \n      }  \n    }\",\n    \"feedbackText\": \"\"\n  },\n  {\n    \"optionText\": \"class Student \n    { \n      int[] scores = new int[5] {3, 2, 4, 1, 5\n      }; \n      public int this[ int index ] \n      {  \n        get \n        {  \n          if (index \u003c 5) return scores[ index ];  \n          else \n          {  \n            Console.WriteLine(\\\"Invalid\\n Index\\\"); return 0;\n            }  \n        }  \n        set \n        {  \n\t  if (index \u003c 5) scores[ index ] = value; else  Console.WriteLine(\\\"Invalid Index\\\");  \n        }  \n      }  \n    }\",\n    \"feedbackText\": \"\"\n  },\n  {\n    \"optionText\": \"class Student { int[] scores = new int[5] {3, 2, 4, 1, 5};  public int this[ int index ] {  get {  if (index \u003c 5) return scores[ index ];  else {  Console.WriteLine(\\\"Invalid Index\\\");  return 0;  }  }  }  }\",\n    \"feedbackText\": \"\"\n  },\n  {\n    \"optionText\": \"class Student { int[] scores = new int[5] {3, 2, 4, 1, 5};  public int this[ int index ] {  get { if (index \u003c 5) scores[ index ] = value;  else {  Console.WriteLine(\\\"Invalid Index\\\"); }  } set {  if (index \u003c 5) return scores[ index ]; else {  Console.WriteLine(\\\"Invalid Index\\\"); return 0; } } } }\",\n    \"feedbackText\": \"\"\n  }]"
	result := modifyString(jsonString)

	// Replace unescaped newline characters with escaped newline characters
	result = strings.ReplaceAll(result, "\n", "\\n")
	result = strings.ReplaceAll(result, "\t", "\\t")

	var quizOptions []QuizOption

	// Unmarshal the JSON string into the struct
	err := json.Unmarshal([]byte(result), &quizOptions)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
	}

	// Display the decoded quiz options
	fmt.Println("Decoded Quiz Options:")
	for _, decodedOption := range quizOptions {
		fmt.Printf("Option Text: %s\nFeedback Text: %s\n\n", decodedOption.OptionText, decodedOption.FeedbackText)
	}

}
