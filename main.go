package main

import (
	"fmt"
	"sync"
	"time"
)

func simpleGoroutine() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Simple print.")
	}()

	wg.Wait()
	fmt.Println("Goroutine finished.")
}

func main() {
	simpleGoroutine()
}
