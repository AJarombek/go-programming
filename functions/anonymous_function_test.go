/**
 * Test file for using anonymous functions in Go
 * Author: Andrew Jarombek
 * Date: 7/4/2022
 */

package functions

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAnonymousFunctions(t *testing.T) {
	// Using the previous pointOp() function as an example, an anonymous function
	// can be defined inline and passed to the function as a parameter.
	point := pointOp(Point{11, 12}, func(i int) int {
		return int(math.Remainder(float64(i), 10))
	})

	assert.Equal(t, 1, point.x)
	assert.Equal(t, 2, point.y)
}
