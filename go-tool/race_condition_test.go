/**
 * Test file for detecting a race condition in concurrent code
 * Author: Andrew Jarombek
 * Date: 8/4/2022
 */

package go_tool

import (
	"testing"
)

// Create a race condition on the 'val' variable between
// two goroutines attempting to increment its value.
var val int

func increment() {
	val++
}

func Increment() {
	for i := 0; i < 1000; i++ {
		increment()
	}
}

func TestRaceCondition(t *testing.T) {
	go Increment()
	go Increment()
}
