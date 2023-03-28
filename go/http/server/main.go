package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const port string = ":9090"
	fmt.Println("http server on port", port)

	mux := http.NewServeMux()

	wrappedMux := middleware(mux)

	mux.HandleFunc("/", homePage)

	server := http.Server{
		Addr:         port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      wrappedMux,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request", r.Host)
	w.Write([]byte("Home Page"))
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UnixMilli()
		handler.ServeHTTP(w, r)
		end := time.Now().UnixMilli()

		fmt.Printf("req time for %s : %v", r.URL.Path, end-start)
	})
}
