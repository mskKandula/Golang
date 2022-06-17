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

	searchResults, err := DataSearch(esclient, ctx, "ABC")
	if err != nil {
		panic(err)
	}

	for _, s := range searchResults {
		fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
	}

}

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func DataSearch(esclient *elastic.Client, ctx context.Context, searchString string) ([]Student, err) {

	var students []Student

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", searchString))

	/* this block will basically print out the es query */
	queryStr, err1 := searchSource.Source()
	if err1 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1)
		return nil, err1
	}

	queryJs, err2 := json.Marshal(queryStr)

	if err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err2)
		return nil, err2
	}

	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	/* until this block */

	searchService := esclient.Search().Index("students").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var student Student
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}
