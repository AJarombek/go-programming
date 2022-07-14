/**
 * Test file for type switches
 * Author: Andrew Jarombek
 * Date: 7/13/2022
 */

package interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeSwitches(t *testing.T) {
	walk := Walk{Miles: 1.76, Description: "To Grand Central for Dad's Birthday"}
	var exercise Exercise = walk

	// Type switch using if-statements
	if exercise == nil {
		assert.Fail(t, "exercise should not be null")
	} else if _, ok := exercise.(ChristmasProduct); ok {
		assert.Fail(t, "exercise should not be a christmas product")
	} else if _, ok = exercise.(Exercise); ok {
		assert.True(t, true)
	}

	// Go provides language support for type switches, making the syntax much easier.
	switch exercise.(type) {
	case nil:
		assert.Fail(t, "exercise should not be null")
	case ChristmasProduct:
		assert.Fail(t, "exercise should not be a christmas product")
	case Exercise:
		assert.True(t, true)
	default:
		assert.Fail(t, "exercise should not reach the default statement")
	}
}
