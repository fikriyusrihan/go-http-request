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
		randWaterInt := h8HelperRand.RandomInt(1, 15)
		randWindInt := h8HelperRand.RandomInt(1, 15)

		requestData := models.StatusRequest{
			Water: uint8(randWaterInt),
			Wind:  uint8(randWindInt),
		}

		requestJson, err := json.MarshalIndent(requestData, "", " ")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(requestJson))

		client := &http.Client{}
		url := "http://localhost:8080"

		req, err := http.NewRequest(
			http.MethodPost,
			url,
			bytes.NewBuffer(requestJson),
		)
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Set("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		responseData := models.StatusResponse{}
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("status water: %s\n", responseData.WaterStatus)
		fmt.Printf("status wind: %s\n\n", responseData.WindStatus)

		_ = res.Body.Close()
		time.Sleep(15 * time.Second)
	}
}
