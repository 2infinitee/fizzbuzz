package main

import (
	"fmt"
	"time"
)

func main() {
	// https://www.callicoder.com/golang-control-flow/

	for i := 0; i < 100; i++ {
		// to slow down output
		time.Sleep(200 * time.Millisecond)

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		}

		if i%3 == 0 {
			fmt.Println("fizz")
		}

		if i%5 == 0 {
			fmt.Println("buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}

	}
}
