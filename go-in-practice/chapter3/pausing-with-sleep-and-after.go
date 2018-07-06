package main

import "time"

func main() {
	// blocks for 5 seconds
	time.Sleep(5 * time.Second)

	// creates a channel that will get notified in 5 seconds, then block until channel receives notification
	sleep := time.After(5 * time.Second)
	<-sleep
}
