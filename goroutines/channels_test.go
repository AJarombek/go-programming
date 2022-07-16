/**
 * Test file for basics of channels in Go
 * Author: Andrew Jarombek
 * Date: 7/15/2022
 */

package goroutines

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Storage struct {
	val int
}

func receiveInt(ch chan int, done chan bool, storage *Storage) {
	// Receive a value through an unbuffered channel
	x := <-ch

	fmt.Printf("Received value %d though a channel.\n", x)
	storage.val = x

	// Send a value through an unbuffered channel,
	// signaling that the function is done processing
	done <- true
}

func TestChannels(t *testing.T) {
	// Create two unbuffered channels
	ch := make(chan int)
	done := make(chan bool)

	storage := Storage{}

	// Run a goroutine, passing both channels as arguments
	go receiveInt(ch, done, &storage)

	// Send a value to the channel
	val := 2
	ch <- val

	// Receive a value from the channel, but ignore its result
	<-done

	assert.Equal(t, val, storage.val)

	// Close the channels.  If channels are not closed, they are garbage collected.
	// Therefore, there is no negative impact to leaving a channel open.
	close(ch)
	close(done)
}
