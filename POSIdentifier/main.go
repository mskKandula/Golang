package main

import (
	"fmt"
	"log"

	prose "github.com/jdkato/prose/v2"
)

func main() {
	// Create a new document with the default configuration:
	doc, err := prose.NewDocument("Go is an open source language supported by Google, developed in Google")
	if err != nil {
		log.Fatal(err)
	}

	set := make(map[string]bool)

	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		if tok.Tag == "NNP" || tok.Tag == "NN" || tok.Tag == "NNS" {
			if _, ok := set[tok.Text]; !ok {
				set[tok.Text] = true
				fmt.Println(tok.Text)
			}

		}
	}

}
