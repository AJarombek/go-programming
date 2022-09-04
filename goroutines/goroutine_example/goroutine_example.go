/**
 * Basic Go code to print text from the main goroutine and another goroutine.
 * Author: Andrew Jarombek
 * Date: 9/2/2022
 */

package main

import (
	"fmt"
	"time"
)

func otherGoroutine() {
	fmt.Println("Other Goroutine")
}

func main() {
	go otherGoroutine()

	// Wait a second in the main goroutine to make it more likely
	// for the other goroutine to run until completion
	time.Sleep(1 * time.Second)

	fmt.Println("Main Goroutine")
}
