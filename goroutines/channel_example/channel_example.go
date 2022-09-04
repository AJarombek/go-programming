/**
 * Basic Go code to pass to data to and from a goroutine via a channel.
 * Author: Andrew Jarombek
 * Date: 9/4/2022
 */

package main

import (
	"fmt"
)

func double(out chan int, in chan int) {
	for {
		value := <-in
		result := value * 2
		out <- result
	}
}

func main() {
	out := make(chan int)
	in := make(chan int)
	go double(out, in)

	in <- 2
	result := <-out
	fmt.Printf("Double 2 = %d\n", result) // Double 2 = 4

	in <- 5
	result = <-out
	fmt.Printf("Double 5 = %d\n", result) // Double 5 = 10

	close(out)
	close(in)
}
