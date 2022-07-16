/**
 * Test file for basics of goroutines in Go
 * Author: Andrew Jarombek
 * Date: 7/13/2022
 */

package goroutines

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func logTime(meta *LogMeta) {
	for {
		fmt.Println(time.Now().Format(time.RFC1123))
		meta.count++
		time.Sleep(time.Second)
	}
}

type LogMeta struct {
	count int
}

func TestBasics(t *testing.T) {
	meta := LogMeta{}

	// logTime runs as a separate goroutine (there is also a main goroutine which is created when a program runs).
	// It's easy to draw parallels between goroutines and threads.  The program begins
	// by running on the main goroutine (thread) and then the "go" keyword creates a
	// separate goroutine (thread) for the logTime() invocation to run within.
	go logTime(&meta)

	// The call to logTime() returns immediately, so wait 3 seconds and let it run.
	time.Sleep(3 * time.Second)

	// This assertion "should" pass
	assert.Equal(t, 3, meta.count)

	// Both the main goroutine and additional goroutines are terminated when the program finishes executing.
}
