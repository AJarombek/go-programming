/**
 * Test file for basics of methods in Go
 * Author: Andrew Jarombek
 * Date: 7/9/2022
 */

package functions

import (
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

// exerciseKilometers is a function converting miles to kilometers for the Exercise type.
func exerciseKilometers(exercise Exercise) float64 {
	return exercise.Miles * 1.609344
}

// exerciseKilometers is a method converting miles to kilometers for the Exercise type.
func (exercise Exercise) exerciseKilometers() float64 {
	return exercise.Miles * 1.609344
}

func TestBasics(t *testing.T) {
	centralParkWalk := Exercise{
		Miles: 2.06,
		Type:  walk,
	}

	km := exerciseKilometers(centralParkWalk)
	assert.Equal(t, 3.3152486400000005, km)

	km = centralParkWalk.exerciseKilometers()
	assert.Equal(t, 3.3152486400000005, km)
}
