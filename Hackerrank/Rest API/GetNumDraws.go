package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

type DataResult struct {
	Competition string `json:"competition"`
	Round       string `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Team1Goals  string `json:"team1goals"`
	Team2Goals  string `json:"team2goals"`
	Year        int32  `json:"year"`
}

type JsonData struct {
	Page       uint8        `json:"page"`
	PerPage    uint8        `json:"per_page"`
	Total      uint32       `json:"total"`
	TotalPages int          `json:"total_pages"`
	Data       []DataResult `json:"data"`
}

func getNumDraws(year int32) int32 {
	var (
		sum      int32
		jsondata JsonData
		wg       sync.WaitGroup
	)
	resp, err := http.Get("https://jsonmock.hackerrank.com/api/football_matches?year=" + strconv.Itoa(int(year)))

	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)

	checkError(err)

	resp.Body.Close()

	err = json.Unmarshal(body, &jsondata)

	checkError(err)

	var i int

	for i = 1; i <= jsondata.TotalPages; i++ {

		wg.Add(1)

		go func(i int) {

			var responseData JsonData

			defer wg.Done()

			resp, err := http.Get("https://jsonmock.hackerrank.com/api/football_matches?year=" + strconv.Itoa(int(year)) + "&page=" + strconv.Itoa(int(i)))

			checkError(err)

			body, err := ioutil.ReadAll(resp.Body)

			checkError(err)

			resp.Body.Close()

			err = json.Unmarshal(body, &responseData)

			checkError(err)

			for _, val := range responseData.Data {
				if val.Team1Goals == val.Team2Goals {
					atomic.AddInt32(&sum, 1)

				}
			}
		}(i)
	}
	wg.Wait()
	return sum
}

func main() {

	var year int32 = 2011

	result := getNumDraws(year)

	fmt.Println(result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
