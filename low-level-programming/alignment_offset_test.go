/**
 * Test file for low-level programming in Go, specifically the alignment of data in memory.
 * Sources: https://stackoverflow.com/q/11386946
 * Author: Andrew Jarombek
 * Date: 8/19/2022
 */

package low_level_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestAlignmentOffset(t *testing.T) {
	// For basic types, the result of unsafe.Alignof() is the same as unsafe.Sizeof().
	stringAlignment := unsafe.Alignof("")
	assert.Equal(t, uintptr(8), stringAlignment)

	boolAlignment := unsafe.Alignof(true)
	assert.Equal(t, uintptr(1), boolAlignment)

	// For complex types like strings, structs and arrays,
	// the result of unsafe.Alignof() is different than unsafe.Sizeof().
	// Alignment specifies where in memory data should be stored,
	// while Sizeof() specifies the size of the data in memory.
	str := "Andy"
	strAlignment := unsafe.Alignof(str)
	strSize := unsafe.Sizeof(str)

	assert.Equal(t, uintptr(8), strAlignment)
	assert.Equal(t, uintptr(16), strSize)

	array := []int{0, 1, 2, 3}
	arrayAlignment := unsafe.Alignof(array)
	arraySize := unsafe.Sizeof(array)

	// For an array, the size in memory is 24 bytes (Sizeof).  Data should be placed
	// in memory at an address which is a multiple of 8 (Alignof).
	assert.Equal(t, uintptr(8), arrayAlignment)
	assert.Equal(t, uintptr(24), arraySize)

	// Performing Sizeof, Alignof, and Offsetof on a struct and its fields
	// unsafe.Offsetof() is the offset in bytes for a field from the start of its structs memory space.
	one := Sample1{a: true, b: true, c: 1.0}

	assert.Equal(t, uintptr(16), unsafe.Sizeof(one))
	assert.Equal(t, uintptr(8), unsafe.Alignof(one))

	assert.Equal(t, uintptr(1), unsafe.Sizeof(one.a))
	assert.Equal(t, uintptr(1), unsafe.Alignof(one.a))
	assert.Equal(t, uintptr(0), unsafe.Offsetof(one.a))

	assert.Equal(t, uintptr(1), unsafe.Sizeof(one.b))
	assert.Equal(t, uintptr(1), unsafe.Alignof(one.b))
	assert.Equal(t, uintptr(1), unsafe.Offsetof(one.b))

	assert.Equal(t, uintptr(8), unsafe.Sizeof(one.c))
	assert.Equal(t, uintptr(8), unsafe.Alignof(one.c))
	assert.Equal(t, uintptr(8), unsafe.Offsetof(one.c))

	two := Sample2{a: true, b: true, c: 1.0}

	assert.Equal(t, uintptr(24), unsafe.Sizeof(two))
	assert.Equal(t, uintptr(8), unsafe.Alignof(two))

	assert.Equal(t, uintptr(1), unsafe.Sizeof(two.a))
	assert.Equal(t, uintptr(1), unsafe.Alignof(two.a))
	assert.Equal(t, uintptr(0), unsafe.Offsetof(two.a))

	assert.Equal(t, uintptr(1), unsafe.Sizeof(two.b))
	assert.Equal(t, uintptr(1), unsafe.Alignof(two.b))
	assert.Equal(t, uintptr(16), unsafe.Offsetof(two.b))

	assert.Equal(t, uintptr(8), unsafe.Sizeof(two.c))
	assert.Equal(t, uintptr(8), unsafe.Alignof(two.c))
	assert.Equal(t, uintptr(8), unsafe.Offsetof(two.c))
}
