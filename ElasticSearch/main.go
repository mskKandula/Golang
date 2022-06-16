package main

import (
	"context"
	"encoding/json"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v7"
)

type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"averageScore"`
}

func main() {
	ctx := context.Background()

	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	//creating student object
	newStudent := Student{
		Name:         "ABC",
		Age:          18,
		AverageScore: 75.5,
	}

	dataJSON, err := json.Marshal(newStudent)
	js := string(dataJSON)
	_, err = esclient.Index().
		Index("students").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
}

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}
