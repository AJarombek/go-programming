/**
 * Test file for using variadic functions in Go
 * Author: Andrew Jarombek
 * Date: 7/7/2022
 */

package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Run struct {
	Miles                   float64
	Hours, Minutes, Seconds int
}

// mergeRuns is a variadic function which takes a variable number of runs
// as an argument and returns a single run.
func mergeRuns(runs ...Run) Run {
	result := Run{}

	for _, run := range runs {
		result.Miles += run.Miles

		seconds := result.Seconds + run.Seconds
		result.Seconds = seconds % 60

		minutes := result.Minutes + run.Minutes + (seconds / 60)
		result.Minutes = minutes % 60

		result.Hours += run.Hours + (minutes / 60)
	}

	return result
}

func TestVariadicFunctions(t *testing.T) {
	shakeout := Run{Miles: 1.51, Hours: 0, Minutes: 11, Seconds: 34}
	warmup := Run{Miles: 2.29, Hours: 0, Minutes: 17, Seconds: 5}
	race := Run{Miles: 4, Hours: 0, Minutes: 21, Seconds: 27}
	cooldown := Run{Miles: 2.24, Hours: 0, Minutes: 16, Seconds: 55}

	july4Runs := mergeRuns(shakeout, warmup, race, cooldown)

	assert.Equal(t, Run{Miles: 10.04, Hours: 1, Minutes: 7, Seconds: 1}, july4Runs)

	// Another way to call mergeRuns()
	runs := []Run{shakeout, warmup, race, cooldown}
	july4Runs = mergeRuns(runs...)

	assert.Equal(t, Run{Miles: 10.04, Hours: 1, Minutes: 7, Seconds: 1}, july4Runs)
}
