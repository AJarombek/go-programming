/**
 * Test file for subtleties when satisfying interfaces
 * Author: Andrew Jarombek
 * Date: 7/10/2022
 */

package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (run *Run) Format(state fmt.State, verb rune) {
	switch verb {
	case 's':
		_, err := fmt.Fprint(state, run.String())

		if err != nil {
			return
		}
	case 'v':
		_, err := fmt.Fprintf(
			state,
			"{Miles=%.02f, Time=%d:%02d:%02d, Feel=%d, Name=%s, Description=%s}",
			run.Miles,
			run.Hours,
			run.Minutes,
			run.Seconds,
			run.Feel,
			run.Name,
			run.Description,
		)

		if err != nil {
			return
		}
	}
}

func TestInterfaceSatisfaction(t *testing.T) {
	// Ensure that Run.Format() works as expected
	run := Run{
		Miles:       4,
		Hours:       0,
		Minutes:     30,
		Seconds:     0,
		Feel:        6,
		Name:        "Marathon Training Block Begins",
		Description: "My planned first run back next Saturday.",
	}

	runPtr := &run
	runStr := fmt.Sprintf("%s", runPtr)
	assert.Equal(t, "Marathon Training Block Begins - 4.00 miles in 0:30:00", runStr)

	runStr = fmt.Sprintf("%v", runPtr)
	assert.Equal(
		t,
		"{Miles=4.00, Time=0:30:00, Feel=6, Name=Marathon Training Block Begins, Description=My planned first run back next Saturday.}",
		runStr,
	)

	// When not using a pointer type, the results are not what I initially expected.
	// Note that the Format() method is not used.  This is because the method is defined with
	// a pointer receiver (*Run), not a value receiver (Run).
	runStr = fmt.Sprintf("%s", run)
	assert.Equal(t, "Marathon Training Block Begins - 4.00 miles in 0:30:00", runStr)

	runStr = fmt.Sprintf("%s", run)
	assert.Equal(t, "Marathon Training Block Begins - 4.00 miles in 0:30:00", runStr)

	// This point is proven further by showing differences in interface satisfaction
	// between Run and *Run.
	var exercise Exercise

	// Type is *Run
	exercise = runPtr

	if _, ok := exercise.(Run); ok {
		assert.Fail(t, "Expected exercise NOT to be of type Run")
	}

	if _, ok := exercise.(Exercise); !ok {
		assert.Fail(t, "Expected exercise to be of type Exercise")
	}

	if _, ok := exercise.(fmt.Stringer); !ok {
		assert.Fail(t, "Expected exercise to be of type fmt.Stringer")
	}

	if _, ok := exercise.(fmt.Formatter); !ok {
		assert.Fail(t, "Expected exercise to be of type fmt.Formatter")
	}

	// Type is Run
	exercise = run

	// Run is of type Run, while *Run is not.
	if _, ok := exercise.(Run); !ok {
		assert.Fail(t, "Expected exercise to be of type Run")
	}

	if _, ok := exercise.(Exercise); !ok {
		assert.Fail(t, "Expected exercise to be of type Exercise")
	}

	if _, ok := exercise.(fmt.Stringer); !ok {
		assert.Fail(t, "Expected exercise to be of type fmt.Stringer")
	}

	// Run is not of type fmt.Formatter, while *Run is.
	if _, ok := exercise.(fmt.Formatter); ok {
		assert.Fail(t, "Expected exercise to NOT be of type fmt.Formatter")
	}
}
