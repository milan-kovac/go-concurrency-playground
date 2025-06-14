package main

import (
	"fmt"
	"net/http"

	simplegoroutine "github.com/milan-kovac/go-concurrency-playground/simple-goroutine"
	worker "github.com/milan-kovac/go-concurrency-playground/worker"
)

func simplegoroutineHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple goroutine")

	simplegoroutine.Run()
}

func workerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Worker")

	worker.Run()
}

func main() {
	http.HandleFunc("/simple-goroutine", simplegoroutineHandler)

	http.HandleFunc("/worker", workerHandler)

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	fmt.Println("Starting server on :8085")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

}
