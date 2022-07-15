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
	go logTime(&meta)
	time.Sleep(3 * time.Second)

	assert.Equal(t, 3, meta.count)
}
