package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://example.com/r/endPoint"
	callPost(url)
}

type Invoice struct {
	InvoiceID string `json:"invoiceId"`
}

type AutoGenerated struct {
	FetchInvoiceByInvoiceID Invoice `json:"FetchInvoiceByInvoiceId"`
}

func callPost(url string) {
	activityName := "FetchInvoiceByInvoiceId"

	jwt := "123.abc.123abc"

	// strData := `{"FetchInvoiceByInvoiceId":{"invoiceId":"guid"}}`

	// strData := `{}`

	// strData, _ = sjson.Set(strData, "invoiceId", "guid")

	// wrapStrData := `{}`

	// wrapStrData, _ = sjson.Set(wrapStrData, "FetchInvoiceByInvoiceId", strData)

	// parsedData := gjson.Parse(wrapStrData)

	// json_data, _ := json.Marshal(parsedData)

	autoGenerated := AutoGenerated{
		FetchInvoiceByInvoiceID: Invoice{
			InvoiceID: "guid",
		},
	}

	// json.Unmarshal(json_data, &autoGenerated)

	jsonData, _ := json.Marshal(autoGenerated)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error while making POST request", err)
		return
	}

	token := "Bearer" + " " + jwt

	req.Header.Set("Service-Header", activityName)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Println("Error while calling POST request", err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println("Error while reading response", err)
		return
	}

	fmt.Println("The response data : ", string(body))
}