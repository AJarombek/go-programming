/**
 * Test file for low-level programming in Go, specifically unsafe pointers and pointer arithmetic.
 * Author: Andrew Jarombek
 * Date: 8/26/2022
 */

package low_level_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestUnsafePointer(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	// Converting a pointer to an unsafe pointer.
	numsPtr := &nums
	numsUnsafePtr := unsafe.Pointer(numsPtr)

	// Converting an unsafe pointer to the data structure it is pointing at.
	nums2 := *(*[]int)(numsUnsafePtr)
	assert.Equal(t, 1, nums[0])
	assert.Equal(t, 1, nums2[0])

	// Perform pointer arithmetic using unsafe pointers.
	// The following code increments a pointer by the size of an integer,
	// therefore pointing to the next element in an array.
	p1 := &nums[0]
	assert.Equal(t, 1, *p1)

	up1 := unsafe.Pointer(p1)
	p2 := (*int)(unsafe.Add(up1, int(unsafe.Sizeof(nums[0]))))
	assert.Equal(t, 2, *p2)

	// Create a slice starting at a specific pointer
	slice := unsafe.Slice(p1, 2)
	assert.Equal(t, 1, slice[0])
	assert.Equal(t, 2, slice[1])
	assert.Equal(t, 2, len(slice))
	assert.Equal(t, 2, cap(slice))
}
