/**
 * Test file for WaitGroups with goroutines and channels in Go.
 * Example is adapted from the one in buffered_channels_test.go.
 * Author: Andrew Jarombek
 * Date: 7/18/2022
 */

package goroutines

import (
	"sync"
	"testing"
)

func makeRequests(endpoints <-chan string) {
	var wg sync.WaitGroup
	counts := make(chan int)
}

func TestWaitGroup(t *testing.T) {

}
