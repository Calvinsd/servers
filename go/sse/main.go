package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/events", handleConnection)

	server := http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	fmt.Println("Starting server at :9090")
	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}

func handleConnection(rw http.ResponseWriter, req *http.Request) {

	fmt.Println("received client connection")

	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "SSE not supported", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("X-Accel-Buffering", "no")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, authorization, X-Requested-With")

	dataChan := make(chan string)

	go streamData(dataChan)

	for data := range dataChan {
		_, err := fmt.Fprint(rw, data)

		if err != nil {
			panic(err)
		}

		fmt.Println("sending data\n", data)
		flusher.Flush()
	}

	defer func() {
		fmt.Println("terminating client connection")
	}()

}

func streamData(dataChan chan<- string) {

	ticker := time.NewTicker(3 * time.Second)
	doneChan := make(chan bool)
	var counter int

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("event: %s\n", "serverevent"))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", "serverdata"))

outerloop:
	for {
		select {
		case <-doneChan:
			break outerloop
		case <-ticker.C:
			{
				dataChan <- sb.String()
				counter++

				if counter > 5 {
					ticker.Stop()
					close(doneChan)
				}
			}
		}
	}

	close(dataChan)
}
