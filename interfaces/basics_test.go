/**
 * Test file for basics of interfaces in Go
 * Author: Andrew Jarombek
 * Date: 7/9/2022
 */

package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Exercise interface {
	fmt.Stringer
}

type Run struct {
	Hours, Minutes, Seconds int
	Miles                   float64
	Feel                    int
	Name                    string
	Description             string
}

// String Because Run implements a String method, it conforms to the Exercise interface
func (run Run) String() string {
	return fmt.Sprintf(
		"%s - %.02f miles in %d:%02d:%02d",
		run.Name,
		run.Miles,
		run.Hours,
		run.Minutes,
		run.Seconds,
	)
}

func TestBasics(t *testing.T) {
	previousLongRun := Run{
		Miles:   12.79,
		Hours:   1,
		Minutes: 36,
		Seconds: 6,
		Name:    "June 19th, Central Park Long Run",
	}

	assert.Equal(t, "June 19th, Central Park Long Run - 12.79 miles in 1:36:06", previousLongRun.String())
}
