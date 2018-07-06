package main

import (
	"time"
	"fmt"
)

func main() {
	// creates a buffered channel with one space
	lock := make(chan bool, 1)
	for i := 1; i < 7; i++ {
		// starts up to 6 goroutines sharing the locking channel
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	// a work acquires the lock by sending it a message
	// the first worker to hit this will get the one space, and thus owns the lock
	// the rest will block
	lock <- true
	fmt.Printf("%d has the lock\n", id)
	// the space between the lock <- true and the <- lock is "locked"
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock\n", id)
	// releases the lock nby reading a value, which then opens that one space
	// on the buffer again so that the next function can lock it
	<-lock
}
