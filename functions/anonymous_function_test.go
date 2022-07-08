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

	// If an anonymous function requires recursion, it must be defined as a variable first.
	var find func(root *TreeNode, val int) bool

	find = func(root *TreeNode, val int) bool {
		if root == nil {
			return false
		}

		if root.Val == val {
			return true
		}

		if root.Val > val {
			return find(root.Left, val)
		} else {
			return find(root.Right, val)
		}
	}

	root := TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	assert.True(t, find(&root, 1))
	assert.True(t, find(&root, 2))
	assert.True(t, find(&root, 3))
	assert.False(t, find(&root, 0))
}
