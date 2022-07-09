/**
 * Test file for Go's defer and recovery mechanisms.
 * Author: Andrew Jarombek
 * Date: 7/9/2022
 */

package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// createRun it is possible (but not recommended) to call panic() when performing
// function validations to handle error scenarios.  panic() causes the program to
// exit with an error code.
func createRun(miles float64, hours, minutes, seconds int) *Run {
	if miles < 0 {
		panic("miles must be greater than or equal to zero")
	}

	if hours < 0 || minutes < 0 || seconds < 0 {
		panic("time must be a positive value")
	}

	if minutes >= 60 || seconds >= 60 {
		panic("minutes and seconds must be between 0 and 59 (inclusive)")
	}

	run := Run{Miles: miles, Hours: hours, Minutes: minutes, Seconds: seconds}
	return &run
}

func TestPanicRecover(t *testing.T) {
	// Panics because miles is less than zero
	assert.Panics(t, func() {
		createRun(-1, 0, 0, 0)
	})

	// Panics because seconds is less than zero
	assert.Panics(t, func() {
		createRun(2, 0, 14, -20)
	})

	// Panics because seconds is greater than 59
	assert.Panics(t, func() {
		createRun(2, 0, 14, 70)
	})

	// Does not panic because the argument values are valid
	assert.NotPanics(t, func() {
		createRun(2, 0, 14, 0)
	})
}
