package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var counterr int64

func Atomic() {
	var wg sync.WaitGroup

	// Atomic values
	fmt.Println("Atomic Values:")
	wg.Add(2)
	go incrementCounterAtomic(&wg)
	go incrementCounterAtomic(&wg)

	wg.Wait()
	fmt.Println("Final Counter Value (Atomic):", atomic.LoadInt64(&counterr))
}

func incrementCounterAtomic(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		atomic.AddInt64(&counterr, 1)
		time.Sleep(time.Millisecond * 100)
	}
}
