/**
 * Test file for buffered channels in Go
 * Author: Andrew Jarombek
 * Date: 7/17/2022
 */

package goroutines

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Links struct {
	Self      string     `json:"self"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Description string `json:"description"`
	Link        string `json:"link"`
	Verb        string `json:"verb"`
}

func request(endpoint string, out chan<- int, err chan<- bool) {
	// TODO
}

func TestBufferedChannels(t *testing.T) {
	// Buffered channel basics
	// Create a buffered channel that can hold five integers.
	ch := make(chan int, 5)

	assert.Equal(t, 0, len(ch))
	assert.Equal(t, 5, cap(ch))

	// Add three elements to the buffered channel.  If the channel was not buffered,
	// "ch <- 2" would hang until another goroutine read the first integer (1) added to the channel.
	ch <- 1
	ch <- 2
	ch <- 3

	assert.Equal(t, 3, len(ch))
	assert.Equal(t, 5, cap(ch))

	// Reading an integer from the channel reduces its length.  In many ways, a channel acts as a queue.
	// However, channels should not be used as a queue in a single goroutine due to the fact they can
	// block the execution of a goroutine.
	assert.Equal(t, 1, <-ch)
	assert.Equal(t, 2, len(ch))

	// Buffered channels example: calling endpoints in api.saintxctf.com to get the total number of endpoints
	// in the API.
	endpoints := []string{
		"https://api.saintsxctf.com/v2/activation_code/links",
		"https://api.saintsxctf.com/v2/comments/links",
		"https://api.saintsxctf.com/v2/flair/links",
		"https://api.saintsxctf.com/v2/forgot_password/links",
		"https://api.saintsxctf.com/v2/logs/links",
		"https://api.saintsxctf.com/v2/log_feed/links",
		"https://api.saintsxctf.com/v2/messages/links",
		"https://api.saintsxctf.com/v2/message_feed/links",
		"https://api.saintsxctf.com/v2/notifications/links",
		"https://api.saintsxctf.com/v2/range_view/links",
		"https://api.saintsxctf.com/v2/teams/links",
		"https://api.saintsxctf.com/v2/users/links",
	}

	counts := make(chan int)
	err := make(chan bool)

	for _, endpoint := range endpoints {
		go func(endpoint string) { request(endpoint, counts, err) }(endpoint)
	}

	result := 0
	for range endpoints {
		result += <-counts
	}

	assert.Equal(t, 0, result)

	success := true
	for range endpoints {
		success = success && <-err
	}

	assert.True(t, success)
}
