/**
 * Test file for using function values in Go
 * Author: Andrew Jarombek
 * Date: 7/4/2022
 */

package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// f A function for increasing the x and y coordinates by a certain exponent power.
// This function isn't called directly; instead it is used in the function values example.
func f(point Point, power int) (x, y int) {
	return int(math.Pow(float64(point.x), float64(power))), int(math.Pow(float64(point.y), float64(power)))
}

// pointOp Function performing an action to a point.
// The action is passed to pointOp as a parameter f.
func pointOp(point Point, f func(int) int) Point {
	return Point{f(point.x), f(point.y)}
}

func TestFunctionValues(t *testing.T) {
	// Functions in Go are first-class values (like languages such as JavaScript, TypeScript).
	// This means functions can be assigned to variables and passed to functions or methods
	// like any other variable.  Functions also have types.
	coordinatesType := fmt.Sprintf("%T", coordinates)
	assert.Equal(t, "func(functions.Point) (int, int)", coordinatesType)

	var powPoint func(point Point, power int) (x, y int)
	assert.Nil(t, powPoint)

	powPoint = f
	assert.NotNil(t, powPoint)

	x, y := powPoint(Point{2, 3}, 3)
	assert.Equal(t, 8, x)
	assert.Equal(t, 27, y)

	// Passing a function as a variable to a function.  This allows dynamic behavior
	// to be passed to a reusable function.
	dub := func(i int) int {
		return i * 2
	}

	point := pointOp(Point{1, 2}, dub)
	assert.Equal(t, 2, point.x)
	assert.Equal(t, 4, point.y)

	log := func(i int) int {
		return int(math.Log2(float64(i)))
	}

	point = pointOp(Point{2, 4}, log)
	assert.Equal(t, 1, point.x)
	assert.Equal(t, 2, point.y)

	log10 := func(i int) int {
		return int(math.Log10(float64(i)))
	}

	point = pointOp(Point{10, 1000}, log10)
	assert.Equal(t, 1, point.x)
	assert.Equal(t, 3, point.y)
}
