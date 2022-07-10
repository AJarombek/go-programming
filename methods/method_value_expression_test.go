/**
 * Test file for basics of method values and method expressions
 * Author: Andrew Jarombek
 * Date: 7/9/2022
 */

package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMethodValueExpression(t *testing.T) {
	harlem5k := FutureExercise{
		Exercise: Exercise{
			Miles: 3.10686,
			Type:  run,
		},
		Date:          "8/13/2022",
		TrainingBlock: "marathon",
	}

	// Create a method value called "kilometersHarlemRace"
	kilometersHarlemRace := harlem5k.exerciseKilometers

	// Invoke the method value.  Notice that no argument is needed.  This seems like
	// a useful way to create partial implementations in Go.
	distance := kilometersHarlemRace()
	assert.Equal(t, 5, int(distance))

	// Method Expression
	kilometers := FutureExercise.exerciseKilometers

	// Creating a method expression is like converting a method to a function.
	distance = kilometers(harlem5k)
	assert.Equal(t, 5, int(distance))
}
