/**
 * Test file for WaitGroups with goroutines and channels in Go.
 * Example is adapted from the one in buffered_channels_test.go.
 * Note: This isn't necessarily best practice, and I do not recommend
 * its use in production code.  However, it does demonstrate the use
 * of sync.WaitGroup for handling an unknown number of goroutines.
 * Author: Andrew Jarombek
 * Date: 7/18/2022
 */

package goroutines

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func makeRequests(endpoints <-chan string) (result int, ok bool) {
	var wg sync.WaitGroup
	var apiCalls int

	counts := make(chan int)
	err := make(chan bool)

	for endpoint := range endpoints {
		wg.Add(1)
		apiCalls++

		go func(endpoint string) {
			defer wg.Done()
			request(endpoint, counts, err)
		}(endpoint)
	}

	go func() {
		wg.Wait()
		close(counts)
		close(err)
	}()

	result = 0
	errorsFound := false

	for i := 0; i < apiCalls; i++ {
		result += <-counts
		errorsFound = errorsFound || <-err
	}

	ok = !errorsFound
	return result, ok
}

func TestWaitGroup(t *testing.T) {
	urls := []string{
		"https://api.saintsxctf.com/v2/activation_code/links",
		"https://api.saintsxctf.com/v2/comments/links",
		"https://api.saintsxctf.com/v2/flair/links",
		"https://api.saintsxctf.com/v2/forgot_password/links",
		"https://api.saintsxctf.com/v2/logs/links",
	}

	endpoints := make(chan string)

	go func(endpoints chan<- string) {
		for i := 0; i <= len(urls); i++ {
			if i < len(urls) {
				endpoints <- urls[i]
			} else {
				close(endpoints)
			}
		}
	}(endpoints)

	result, ok := makeRequests(endpoints)

	assert.Equal(t, 20, result)
	assert.True(t, ok)
}
