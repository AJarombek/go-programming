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

func f(point Point, power int) (x, y int) {
	return int(math.Pow(float64(point.x), float64(power))), int(math.Pow(float64(point.y), float64(power)))
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
}
