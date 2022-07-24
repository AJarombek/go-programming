/**
 * Techniques for limiting the number of concurrent goroutines.
 * Author: Andrew Jarombek
 * Date: 7/23/2022
 */

package goroutines

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

// numberSeries an unnecessarily complex implementation to find the nth result in the sequence 1 + 2 + 3 + 4...
// Demonstrates how to limit the number of concurrent goroutines.
func numberSeries(x int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var result int

	values := make(chan int, 1_024)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			for value := range values {
				mu.Lock()
				result += value
				println(result)
				mu.Unlock()
			}

			wg.Done()
		}()
	}

	for i := 0; i < x; i++ {
		values <- i + 1
	}

	close(values)
	wg.Wait()

	return result
}

func TestLimitGoroutines(t *testing.T) {
	result := numberSeries(10)
	assert.Equal(t, 55, result)

	result = numberSeries(20)
	assert.Equal(t, 210, result)
}
