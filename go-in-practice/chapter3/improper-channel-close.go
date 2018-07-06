package main

import (
	"time"
	"fmt"
)

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)

	// start goroutine with sending channel
	go send(msg)

	// loop over a select that watches for messages from send, or for a time-out
	for {
		select {
			case m := <-msg:
				// message arrives from send, print it
				fmt.Println(m)
			case <-until:
				// when the timeout occurs, shuts things down
				// pauses to ensure that you see failure before main goroutine exits
				close(msg)
				time.Sleep(500 * time.Millisecond)
				return
		}
	}
}

func send(ch chan string) {
	// sends hello to channel every half second
	for {
		ch <- "hello"
		time.Sleep(500 * time.Millisecond)
	}
}