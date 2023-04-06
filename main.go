package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"water-wind-client/models"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

func main() {
	for true {
		// Initialize random number and request data
		randWaterInt := h8HelperRand.RandomInt(1, 15)
		randWindInt := h8HelperRand.RandomInt(1, 15)

		requestData := models.Status{
			Water: uint8(randWaterInt),
			Wind:  uint8(randWindInt),
		}

		requestJson, err := json.MarshalIndent(requestData, "", " ")
		if err != nil {
			log.Fatalln(err)
		}

		// Prepare request to server
		client := &http.Client{}
		url := "https://jsonplaceholder.typicode.com/posts"

		req, err := http.NewRequest(
			http.MethodPost,
			url,
			bytes.NewBuffer(requestJson),
		)
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Set("Content-Type", "application/json")

		// Send request to server
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// Parse response from server
		var responseData models.Status
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			log.Fatalln(err)
		}

		responseJson, err := json.MarshalIndent(responseData, "", " ")
		if err != nil {
			log.Fatalln(err)
		}

		// Print response in json format
		fmt.Println(string(responseJson))

		// Determine status
		if requestData.Water <= 5 {
			fmt.Printf("status water: aman\n")
		} else if requestData.Water >= 6 && requestData.Water <= 8 {
			fmt.Printf("status water: siaga\n")
		} else {
			fmt.Printf("status water: bahaya\n")
		}

		if requestData.Wind <= 6 {
			fmt.Printf("status wind: aman\n")
		} else if requestData.Wind >= 7 && requestData.Wind <= 15 {
			fmt.Printf("status wind: siaga\n")
		} else {
			fmt.Printf("status wind: bahaya\n")
		}

		_ = res.Body.Close()
		time.Sleep(15 * time.Second)
	}
}
