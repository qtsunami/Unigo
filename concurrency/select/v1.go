package main

import (
	"time"
	"fmt"
)

func main() {
	startTime := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")

	select {
	case <-c:
		fmt.Printf("Unblocked  %v later", time.Since(startTime))
	}
}