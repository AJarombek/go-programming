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
	var coreType reflect.Type = reflect.TypeOf(core)
	assert.Equal(t, "reflection.Core", coreType.String())

	// Gain insights into the type using methods of reflect.Type
	assert.Equal(t, 1, coreType.NumMethod())
	assert.Equal(t, 2, coreType.NumField())

	assert.Equal(t, "time", coreType.Field(0).Name)
	assert.Equal(t, "github.com/ajarombek/go-programming/reflection", coreType.Field(0).PkgPath)

	assert.Equal(t, "date", coreType.Field(1).Name)
	assert.Equal(t, "github.com/ajarombek/go-programming/reflection", coreType.Field(1).PkgPath)

	assert.Equal(t, "Summary", coreType.Method(0).Name)

	// Ensure the Core type implements Exercise
	exerciseType := reflect.TypeOf((*Exercise)(nil)).Elem()
	assert.True(t, coreType.Implements(exerciseType))

	// Ensure the Core type implements interface{}
	interfaceType := reflect.TypeOf((*interface{})(nil)).Elem()
	assert.True(t, coreType.Implements(interfaceType))

	// Ensure the Core type DOES NOT implement Stringer
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	assert.False(t, coreType.Implements(stringerType))

	// Behind the scenes, "%T" in formatted strings uses reflection to determine types
	assert.Equal(t, "reflection.Core", fmt.Sprintf("%T", core))

	// Hold the value of a type using reflect.ValueOf()
	v := reflect.ValueOf("Jarombek")

	assert.Equal(t, "Jarombek", v.String())
	assert.Equal(t, reflect.String, v.Kind())

	// "%v" in formatted strings prints the value of reflect.Value types
	assert.Equal(t, "Jarombek", fmt.Sprintf("%v", v))

	var coreValue reflect.Value = reflect.ValueOf(core)

	// Kind() can be used to group values into general type categories
	assert.Equal(t, reflect.Struct, coreValue.Kind())

	// Call a method using reflection
	resultValues := coreValue.MethodByName("Summary").Call([]reflect.Value{})
	result := resultValues[0].Interface()
	assert.Equal(t, "Core Exercise: 22:25 on Aug 11 2022 10:10 PM", result)

	// Another way to call the method
	resultValues = coreValue.Method(0).Call([]reflect.Value{})
	result = resultValues[0].Interface()

	// result is an interface{}
	assert.Equal(t, "Core Exercise: 22:25 on Aug 11 2022 10:10 PM", result)

	// result.(string) is a string
	assert.Equal(t, "Core Exercise: 22:25 on Aug 11 2022 10:10 PM", result.(string))

	// Alter a value using reflection
	name := "Andrew"
	elem := reflect.ValueOf(&name).Elem()

	// Assert that 'elem' is addressable - meaning it has a
	// storage location (address) on the stack or heap.  If the
	// reflect.Value variable isn't addressable, its value can't be changed.
	assert.True(t, elem.CanAddr())
	assert.True(t, elem.CanSet())

	ptrName := elem.Addr().Interface().(*string)
	*ptrName = "Andy"
	assert.Equal(t, "Andy", name)

	/*
		All together, the Go equivalent of:
			name = "Andy"
		in reflection is:
			*reflect.ValueOf(&name).Elem().Addr().Interface().(*string) = "Andy"
	*/

	*reflect.ValueOf(&name).Elem().Addr().Interface().(*string) = "Andy Jarombek"
	assert.Equal(t, "Andy Jarombek", name)

	// Another way to alter a value using reflection
	reflect.ValueOf(&name).Elem().Set(reflect.ValueOf("Andrew Jarombek"))
	assert.Equal(t, "Andrew Jarombek", name)

	// Finally, one more way to alter a value using reflection
	reflect.ValueOf(&name).Elem().SetString("Andrew")
	assert.Equal(t, "Andrew", name)
}
