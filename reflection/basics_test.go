/**
 * Test file for the basics of reflectio in Go
 * Author: Andrew Jarombek
 * Date: 8/13/2022
 */

package reflection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

type Exercise interface {
	Summary() string
}

type Core struct {
	time string
	date time.Time
}

func (core Core) Summary() string {
	return fmt.Sprintf("Core Exercise: %s on %s", core.time, core.date.Format("Jan 2 2006 3:04 PM"))
}

func TestBasics(t *testing.T) {
	strType := reflect.TypeOf("Andy")

	assert.Equal(t, "string", strType.String())
	assert.Equal(t, "string", strType.Name())

	location, _ := time.LoadLocation("America/New_York")
	var core Exercise = Core{
		time: "22:25",
		date: time.Date(2022, 8, 11, 22, 10, 0, 0, location),
	}

	assert.Equal(t, "Core Exercise: 22:25 on Aug 11 2022 10:10 PM", core.Summary())

	// Reflection type for an interface is the concrete type (reflection.Core),
	// not the interface type (reflection.Exercise)
	coreType := reflect.TypeOf(core)
	assert.Equal(t, "reflection.Core", coreType.String())
}
