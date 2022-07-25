/**
 * Test file for canceling goroutines using channels in Go
 * Author: Andrew Jarombek
 * Date: 7/24/2022
 */

package goroutines

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func canceled(done <-chan struct{}) bool {
	// When a select statement has a "default" case, it acts as a polling
	// operator that doesn't wait for a value to enter a channel.  If all
	// the case channels are empty, the "default" case is used.
	select {
	case <-done:
		return true
	default:
		return false
	}
}

type Count struct {
	i int
}

func count(done <-chan struct{}, c *Count) {
	for {
		if canceled(done) {
			return
		}

		ticker := time.NewTicker(time.Second)

		<-ticker.C
		c.i++
		fmt.Println("Add 1")
	}
}

func TestCancelGoroutine(t *testing.T) {
	done := make(chan struct{})
	c := Count{}

	go count(done, &c)

	time.Sleep(3 * time.Second)
	done <- struct{}{}
	fmt.Printf("Canceled")

	time.Sleep(3 * time.Second)

	// It is still possible this assertion fails, although it is unlikely.
	assert.Equal(t, 3, c.i)
}
