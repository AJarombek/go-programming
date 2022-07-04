/**
 * Test file for creating JSON from structs and converting structs to JSON
 * Author: Andrew Jarombek
 * Date: 7/1/2022
 */

package functions

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Point struct {
	x int
	y int
}

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

// describeTree is a recursive function that generates a list of strings with every root->leaf path
// in a binary tree.
func describeTree(nodes []int, root *TreeNode) []string {
	nodes = append(nodes, root.Val)
	result := []string{fmt.Sprint(nodes)}

	if root.Left != nil {
		left := fmt.Sprint(describeTree(nodes, root.Left))
		result = append(result, left[1:len(left)-1])
	}

	if root.Right != nil {
		right := fmt.Sprint(describeTree(nodes, root.Right))
		result = append(result, right[1:len(right)-1])
	}

	return result
}

// coordinates Deconstruct the x and y coordinates in a point.  Very basic example of a
// function returning multiple values.
func coordinates(p Point) (int, int) {
	return p.x, p.y
}

// sqrtPoint Performs math.Sqrt() o both coordinates, returning both as separate values.
// This function demonstrates how to assign names to return types.
func sqrtPoint(p Point) (x float64, y float64) {
	return math.Sqrt(float64(p.x)), math.Sqrt(float64(p.y))
}

// createExerciseLog creates an exercise log from multiple inputs.  It uses the concept
// of "bare returns", which are possible with named return types.  It also demonstrates
// how consecutive parameters with the same return type only require an explicit type
// declaration after the last parameter.
func createExerciseLog(hours, minutes, seconds, feel int, miles float64, name, description, exerciseType string) (log Exercise, err error) {
	if hours < 0 || minutes < 0 || seconds < 0 || minutes > 59 || seconds > 59 {
		err = errors.New("time out of bounds")
		return
	}

	if feel <= 0 || feel > 10 {
		err = errors.New("feel must be between 1 and 10 inclusive")
		return
	}

	// TODO
	return
}

func TestBasics(t *testing.T) {
	// Use recursion and the nature of slices to generate and print a mapping of the tree node paths.
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	result := describeTree(nil, &root)

	assert.Equal(t, 3, len(result))

	fmt.Println(result)
	assert.Equal(t, "[1]", result[0])
	assert.Equal(t, "[1 2]", result[1])
	assert.Equal(t, "[1 3]", result[2])

	// Functions in Go commonly have multiple return values
	x, y := coordinates(Point{1, 2})
	assert.Equal(t, 1, x)
	assert.Equal(t, 2, y)

	sqrtX, sqrtY := sqrtPoint(Point{4, 5})
	assert.Equal(t, 2.0, sqrtX)
	assert.Equal(t, 2.23606797749979, sqrtY)
}
