package main

import "fmt"

// pipeline

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	// squares
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// printer
	for x := range squares {
		fmt.Println(x)
	}

}
