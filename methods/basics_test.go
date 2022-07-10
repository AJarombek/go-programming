/**
 * Test file for basics of methods in Go
 * Author: Andrew Jarombek
 * Date: 7/9/2022
 */

package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	run  = "run"
	bike = "bike"
	swim = "swim"
	walk = "walk"
)

type Exercise struct {
	Hours, Minutes, Seconds int
	Miles                   float64
	Feel                    int
	Name                    string
	Description             string
	Type                    string
}

type FutureExercise struct {
	Exercise
	Date          string
	TrainingBlock string
}

// exerciseKilometers is a function converting miles to kilometers for the Exercise type.
func exerciseKilometers(exercise Exercise) float64 {
	return exercise.Miles * 1.609344
}

// exerciseKilometers is a method converting miles to kilometers for the Exercise type.
func (exercise Exercise) exerciseKilometers() float64 {
	return exercise.Miles * 1.609344
}

// timeString is a method with a pointer receiver (*Exercise) that creates a
// string representation of the time spent exercising.  In general, it's recommended
// to have all methods for a type be either value receivers or pointer receivers, not
// a mix of both like in this file.
func (exercise *Exercise) timeString() string {
	return fmt.Sprintf("%d:%02d:%02d", exercise.Hours, exercise.Minutes, exercise.Seconds)
}

func TestBasics(t *testing.T) {
	centralParkWalk := Exercise{
		Miles:   2.06,
		Type:    walk,
		Hours:   0,
		Minutes: 38,
		Seconds: 4,
	}

	km := exerciseKilometers(centralParkWalk)
	assert.Equal(t, 3.3152486400000005, km)

	km = centralParkWalk.exerciseKilometers()
	assert.Equal(t, 3.3152486400000005, km)

	// Call a method with a pointer receiver
	time := (&centralParkWalk).timeString()
	assert.Equal(t, "0:38:04", time)

	// Shorthand call with an implicit conversion from "centralParkWalk" to "(&centralParkWalk)"
	time = centralParkWalk.timeString()
	assert.Equal(t, "0:38:04", time)

	// Call a method on a pointer with a pointer receiver
	walkPtr := &centralParkWalk
	time = (*walkPtr).timeString()
	assert.Equal(t, "0:38:04", time)

	// Shorthand call with an implicit conversion from "walkPtr" to "(*walkPtr)"
	time = walkPtr.timeString()
	assert.Equal(t, "0:38:04", time)

	// Composing types through struct embedding
	teamChamps := FutureExercise{
		Exercise: Exercise{
			Miles: 5,
			Type:  run,
		},
		Date:          "7/31/2022",
		TrainingBlock: "marathon",
	}

	// Composing types creates a "has a" relationship.  In this case, "FutureExercise has an Exercise", meaning
	// FutureExercise has all the fields of Exercise and can use its methods.  In the following example, a value
	// of type FutureExercise uses the exerciseKilometers() method, which is defined on Exercise.
	km = teamChamps.exerciseKilometers()
	assert.Equal(t, 8.04672, km)

	// Equivalent (but longer) method call
	km = teamChamps.Exercise.exerciseKilometers()
	assert.Equal(t, 8.04672, km)
}
