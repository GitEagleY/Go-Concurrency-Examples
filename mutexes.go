package main

import (
	"fmt"
	"sync"
	"time"
)

var counterrr int
var muu sync.Mutex

func Mutexes() {
	var wg sync.WaitGroup

	// Mutexes
	fmt.Println("\nMutexes:")
	wg.Add(2)
	go incrementCounterMutex(&wg)
	go incrementCounterMutex(&wg)

	wg.Wait()
	fmt.Println("Final Counter Value (Mutex):", counterrr)
}

func incrementCounterMutex(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		muu.Lock()
		counterrr++
		muu.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
}
