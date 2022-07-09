/**
 * Test file for Go's defer and recovery mechanisms.
 * Author: Andrew Jarombek
 * Date: 7/9/2022
 */

package functions

import (
	"errors"
	"fmt"
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

// createRunV2 is the same as createRun except it uses the preferred approach for
// handling error conditions in Go.  Instead of panicking, it returns an error value
// along with a run value for the caller to handle.
func createRunV2(miles float64, hours, minutes, seconds int) (err error, run *Run) {
	if miles < 0 {
		err = errors.New("miles must be greater than or equal to zero")
		return err, nil
	}

	if hours < 0 || minutes < 0 || seconds < 0 {
		err = errors.New("time must be a positive value")
		return err, nil
	}

	if minutes >= 60 || seconds >= 60 {
		err = errors.New("minutes and seconds must be between 0 and 59 (inclusive)")
		return err, nil
	}

	run = &Run{Miles: miles, Hours: hours, Minutes: minutes, Seconds: seconds}
	return nil, run
}

// panicRecover demonstrates how to recover from a panic without exiting the program with an error code.
// It utilizes the recover() function to avoid a panic.
func panicRecover() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("recovered from a panic")
		}
	}()

	panic("panicking")
}

// conditionalPanicRecover attempts to recover from a panic depending on the type of panic encountered.
func conditionalPanicRecover(i int) (err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
			return
		case "expected panic":
			err = fmt.Errorf("expected panic encountered")
		default:
			panic(p)
		}
	}()

	if i%3 == 0 {
		panic("expected panic")
	} else if i%3 == 1 {
		panic("unexpected panic")
	}

	return nil
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

	// Preferred approach to panicking
	err, _ := createRunV2(-1, 0, 0, 0)
	assert.Error(t, err)

	err, run := createRunV2(2, 0, 14, 0)
	assert.Nil(t, err)
	assert.Equal(t, 2.0, run.Miles)

	assert.NotPanics(t, func() {
		err := panicRecover()
		if err == nil {
			panic("expected an error")
		}
	})

	assert.NotPanics(t, func() {
		err := conditionalPanicRecover(0)
		if err == nil {
			panic("expected an error")
		}
	})

	assert.Panics(t, func() {
		err := conditionalPanicRecover(1)
		if err != nil {
			panic("expected a panic")
		}
	})

	assert.NotPanics(t, func() {
		err := conditionalPanicRecover(2)
		if err != nil {
			panic("expected a successful response")
		}
	})
}
