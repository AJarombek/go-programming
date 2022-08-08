/**
 * Unit tests for current_times program.
 * Author: Andrew Jarombek
 * Date: 8/7/2022
 */

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrentTime(t *testing.T) {
	time, err := currentTime("America/New_York")
	assert.Nil(t, err)
	assert.Greater(t, len(time), 0)

	time, err = currentTime("America/Stamford")
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(time))
}
