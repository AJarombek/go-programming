/**
 * Test file for basics of interfaces in Go
 * Author: Andrew Jarombek
 * Date: 7/9/2022
 */

package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Exercise interface {
	fmt.Stringer
}

type Run struct {
	Hours, Minutes, Seconds int
	Miles                   float64
	Feel                    int
	Name                    string
	Description             string
}

// String Because Run implements a String method, it conforms to the Exercise interface
func (run Run) String() string {
	return fmt.Sprintf(
		"%s - %.02f miles in %d:%02d:%02d",
		run.Name,
		run.Miles,
		run.Hours,
		run.Minutes,
		run.Seconds,
	)
}

func TestBasics(t *testing.T) {
	previousLongRun := Run{
		Miles:   12.79,
		Hours:   1,
		Minutes: 36,
		Seconds: 6,
		Name:    "June 19th, Central Park Long Run",
	}

	assert.Equal(t, "June 19th, Central Park Long Run - 12.79 miles in 1:36:06", previousLongRun.String())

	// Since the Run type satisfies the Exercise interface, a value of type Run can be assigned
	// to a variable of type Exercise.  The same holds true for assigning a value of type Run to
	// a variable of type fmt.Stringer.
	var exercise Exercise
	exercise = previousLongRun

	if _, ok := exercise.(Run); !ok {
		assert.Fail(t, "Expected exercise to be of type Run")
	}

	if _, ok := exercise.(Exercise); !ok {
		assert.Fail(t, "Expected exercise to be of type Exercise")
	}

	if _, ok := exercise.(fmt.Stringer); !ok {
		assert.Fail(t, "Expected exercise to be of type fmt.Stringer")
	}

	if _, ok := exercise.(fmt.Formatter); ok {
		assert.Fail(t, "Expected exercise NOT to be of type fmt.Formatter")
	}

	var stringer fmt.Stringer
	stringer = previousLongRun

	if _, ok := stringer.(Run); !ok {
		assert.Fail(t, "Expected stringer to be of type Run")
	}

	if _, ok := stringer.(Exercise); !ok {
		assert.Fail(t, "Expected stringer to be of type Exercise")
	}

	if _, ok := stringer.(fmt.Stringer); !ok {
		assert.Fail(t, "Expected stringer to be of type fmt.Stringer")
	}

	if _, ok := stringer.(fmt.Formatter); ok {
		assert.Fail(t, "Expected stringer NOT to be of type fmt.Formatter")
	}

	// This would result in a compile time failure, so it must be commented out.
	// The fmt.Formatter interface requires types to have a Format() method.

	// var formatter fmt.Formatter
	// formatter = previousLongRun

	// It can also be assigned to an empty interface
	var emptyInterface interface{}
	emptyInterface = previousLongRun

	if _, ok := emptyInterface.(Run); !ok {
		assert.Fail(t, "Expected emptyInterface to be of type Run")
	}

	if _, ok := emptyInterface.(Exercise); !ok {
		assert.Fail(t, "Expected emptyInterface to be of type Exercise")
	}

	if _, ok := emptyInterface.(fmt.Stringer); !ok {
		assert.Fail(t, "Expected emptyInterface to be of type fmt.Stringer")
	}

	if _, ok := emptyInterface.(fmt.Formatter); ok {
		assert.Fail(t, "Expected emptyInterface NOT to be of type fmt.Formatter")
	}

	if _, ok := emptyInterface.(interface{}); !ok {
		assert.Fail(t, "Expected emptyInterface to be of type empty interface")
	}

	// Example continues in interface_satisfaction_test.go
}
