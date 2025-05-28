package simplegoroutine

import (
	"fmt"
	"sync"
	"time"
)

func runSimpleGoroutine() {
	var wg sync.WaitGroup

	// increases the number of goroutines we need to wait for
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Simple print.")
	}()

	// blocks the main thread
	wg.Wait()
	fmt.Println("Goroutine finished.")
}

func runGoroutineWithChannel() {

	// Creates an unbuffered channel that transmits only strings
	ch := make(chan string)

	// Creates a buffered channel that transmits only strings and can hold up to n values without blocking
	// ch := make(chan string, n)

	go func() {
		// Send 5 messages into the channel
		for i := 1; i <= 5; i++ {
			ch <- fmt.Sprintf("Message %d", i)
			time.Sleep(100 * time.Millisecond)
		}
		// Close the channel after sending all messages to signal no more data will come
		close(ch)
	}()

	// Loops over messages received from the channel until it is closed
	for msg := range ch {
		fmt.Println(msg)
	}
}

func Run() {
	// without go this takes place in the main thread
	runSimpleGoroutine()

	// go starts a goroutine — a lightweight concurrent unit that is NOT a real OS thread, but managed by the Go runtime on OS threads.
	go runGoroutineWithChannel()
	go runGoroutineWithChannel()

	fmt.Println("Simple example run completed.")
}
