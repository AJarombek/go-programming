/**
 * Test file for using the select statement for multiplexing (receiving and handling data from multiple channels).
 * Author: Andrew Jarombek
 * Date: 7/23/2022
 */

package goroutines

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func op(x int, y int, add chan struct{}, sub chan struct{}, mul chan struct{}, div chan struct{}, end chan struct{}, result chan int) {
	for {
		select {
		case <-add:
			result <- x + y
		case <-sub:
			result <- x - y
		case <-mul:
			result <- x * y
		case <-div:
			result <- x / y
		case <-end:
			return
		}
	}
}

func TestMultiplexingSelect(t *testing.T) {
	add := make(chan struct{})
	sub := make(chan struct{})
	mul := make(chan struct{})
	div := make(chan struct{})
	end := make(chan struct{})
	result := make(chan int)

	go op(2, 3, add, sub, mul, div, end, result)

	add <- struct{}{}
	assert.Equal(t, 5, <-result)

	sub <- struct{}{}
	assert.Equal(t, -1, <-result)

	mul <- struct{}{}
	assert.Equal(t, 6, <-result)

	div <- struct{}{}
	assert.Equal(t, 0, <-result)

	end <- struct{}{}

	close(add)
	close(sub)
	close(mul)
	close(div)
	close(result)
}
