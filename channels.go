package main

import (
	"fmt"
	"sync"
	"time"
)

func Channels() {
	var wg sync.WaitGroup

	// Channels
	fmt.Println("\nChannels:")
	ch := make(chan string, 2)

	wg.Add(2)
	go sendData(ch, &wg)
	go receiveData(ch, &wg)

	wg.Wait()
	close(ch)
}

func sendData(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Message %d", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func receiveData(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}
