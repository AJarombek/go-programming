/**
 * Test file for the "defer" keyword.
 * Author: Andrew Jarombek
 * Date: 7/9/2022
 */

package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// increment increases an integer by 1.
func increment(x int) int {
	return x + 1
}

// double shows some limitations of the "defer" keyword.
// The value of the return statement will not be impacted by "defer",
// unless the value is a reference type.
func double(x int) int {
	defer increment(x)
	x *= 2
	return x
}

// incrementPoint increases the x and y coordinates in a point by 1.
func incrementPoint(p *Point) {
	p.y++
	p.x++
}

// doubleIncrement shows that defer impacts the returned value when it is a pointer.
// It is important to note that although defined prior to the multiplication operations
// in the function lexical scope, the defer command is invoked afterwards.
func doubleIncrement(p *Point) *Point {
	defer incrementPoint(p)

	p.y *= 2
	p.x *= 2

	return p
}

func TestDefer(t *testing.T) {
	assert.Equal(t, 4, double(2))

	point := Point{2, 4}
	point = *doubleIncrement(&point)

	assert.Equal(t, 5, point.x)
	assert.Equal(t, 9, point.y)
}
