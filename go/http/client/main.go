package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Client running")

	const url string = "http://localhost:9090"
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error in making request")
		panic(err)
	}

	defer res.Body.Close()

	var body []byte

	res.Body.Read(body)

	fmt.Println(res)
	fmt.Println(body)
}
