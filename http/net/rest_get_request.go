package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
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
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod(fasthttp.MethodGet)
	request.Header.SetRequestURI(url)
	request.Header.Add("Accept", "application/json")
	defer fasthttp.ReleaseRequest(request) //at the end of the method, release the request

	response := fasthttp.AcquireResponse() //send request to pool, if the response cannot be found at the pool, then creates a new response
	defer fasthttp.ReleaseResponse(response)

	err := fasthttp.Do(request, response)
	if err != nil {
		return nil, err
	}

	return response.Body(), nil
}
