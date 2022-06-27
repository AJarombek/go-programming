/**
 * Main test file for composite types
 * Author: Andrew Jarombek
 * Date: 6/27/2022
 */

package _go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Exercise struct {
	Hours, Minutes, Seconds uint8
	Miles float64
	Feel uint8
	Name string
	Description string
}

type Run struct {
	Exercise
	Shoe string
}

type Bike struct {
	Exercise
	Model string
}

func Test(t *testing.T) {
	// Struct embedding allows nested fields to be assigned to and accessed
	// on the top level struct variable.
	var easyRun Run
	easyRun.Hours = 0
	easyRun.Minutes = 56
	easyRun.Seconds = 50
	easyRun.Miles = 7.66
	easyRun.Feel = 6
	easyRun.Name = "Cognewaugh & Cat Rock"
	easyRun.Description = "Tired from lots of driving, saw some deer which made the run worthwhile"
	easyRun.Shoe = "Asics GT-2000"

	assert.Equal(t, uint8(56), easyRun.Minutes)
	assert.Equal(t, easyRun.Exercise.Name, easyRun.Name)

	// Another way to create an embedded struct
	hillyRun := Run{
		Exercise: Exercise{
			Hours: 0,
			Minutes: 46,
			Seconds: 33,
			Miles: 5.64,
			Feel: 4,
			Name: "Exploring SF",
			Description: "Late flight so only got 4 hours of sleep, very tired.  Tech conference starts tomorrow",
		},
		Shoe: "Asics GT-2000",
	}

	assert.Equal(t, "Exploring SF", hillyRun.Name)
}