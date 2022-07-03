/**
 * Test file for creating JSON from structs and converting structs to JSON
 * Author: Andrew Jarombek
 * Date: 7/1/2022
 */

package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func TestBasics(t *testing.T) {
	// Use recursion and the nature of slices to generate and print a mapping of the tree node paths.
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	result := describeTree(nil, &root)

	assert.Equal(t, 3, len(result))

	fmt.Println(result)
	assert.Equal(t, "[1]", result[0])
	assert.Equal(t, "[1 2]", result[1])
	assert.Equal(t, "[1 3]", result[2])
}
