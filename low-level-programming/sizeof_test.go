/**
 * Test file for low-level programming in Go, specifically the unsafe.Sizeof() function.
 * Author: Andrew Jarombek
 * Date: 8/19/2022
 */

package low_level_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

type Sample1 struct {
	a bool
	b bool
	c float64
}

type Sample2 struct {
	a bool
	c float64
	b bool
}

func TestSizeOf(t *testing.T) {
	// unsafe.Sizeof() gets the size of data in bytes.  unsafe.Sizeof() is not necessarily 'unsafe',
	// but its resulting values will depend on the architecture of the machine it is run on.
	sizeOfInt := unsafe.Sizeof(int(0))

	// On a 32-bit machine, an int would be 4 bytes instead of 8.  Since this repository runs this
	// code on a Docker container, I can feel confident that the result will be 8 since the architecture
	// of the container ~shouldn't~ change.
	assert.Equal(t, uintptr(8), sizeOfInt)

	sizeOfFloat64 := unsafe.Sizeof(float64(0))
	assert.Equal(t, uintptr(8), sizeOfFloat64)

	sizeOfString := unsafe.Sizeof("Andy")
	assert.Equal(t, uintptr(16), sizeOfString)

	andy := "Andy"
	sizeOfPtr := unsafe.Sizeof(&andy)
	assert.Equal(t, uintptr(8), sizeOfPtr)

	name := []string{"Andy", "Jarombek"}
	sizeOfArray := unsafe.Sizeof(name)

	// Size of an array is 24 bytes because 8 bytes are used to track the length,
	// 8 bytes are used to track the capacity, and 8 bytes are for the data.
	// NOTE: The bytes used to represent the data will always be 8 bytes (on the current architecture),
	// despite the length of the array.  You can think of the 8 bytes as being a
	// pointer to the location in memory holding the array elements.  The same is true
	// for strings - unsafe.Sizeof() will always return the same value on the same machine
	// no matter the length of the string.
	assert.Equal(t, uintptr(24), sizeOfArray)

	sizeOfBool := unsafe.Sizeof(false)
	assert.Equal(t, uintptr(1), sizeOfBool)

	// The size of structs in memory can depend on the order the fields are defined.
	// For example, Sample1 and Sample2 have the same fields and appear to be the same,
	// but their sizes in memory are different.  This is due to computers aligning data in memory,
	// and the creation of gaps in memory if an item takes up less bytes than its word size
	// (8 bytes in a 64-bit machine, 4 bytes in a 32-bit machine).
	one := Sample1{a: true, b: true, c: 1.0}
	two := Sample2{a: true, b: true, c: 1.0}

	assert.Equal(t, uintptr(16), unsafe.Sizeof(one))
	assert.Equal(t, uintptr(24), unsafe.Sizeof(two))
}
