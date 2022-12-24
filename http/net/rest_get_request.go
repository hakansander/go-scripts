package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const url = "https://icanhazdadjoke.com/"

type joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	start := time.Now()
	response, err := getHttp()
	if err != nil {
		log.Fatal(err)
	}
	elapsedTime := time.Since(start)

	var joke joke
	if err := json.Unmarshal(response, &joke); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Elapsed time is: %f seconds, Data is: %v\n", elapsedTime.Seconds(), joke)
}

func getHttp() ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseByte))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	return responseByte, nil
}
