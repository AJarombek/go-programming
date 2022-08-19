/**
 * Test file for using maps with the reflection API in Go
 * Author: Andrew Jarombek
 * Date: 8/19/2022
 */

package reflection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMapReflection(t *testing.T) {
	// Top 7 languages I've written code for in 2022 (so far)
	usage2022 := make(map[string]int)
	usage2022["Python"] = 10_319
	usage2022["Go"] = 5_482
	usage2022["TypeScript"] = 4_760
	usage2022["C"] = 3_369
	usage2022["HTML"] = 1_781
	usage2022["Java"] = 1_563
	usage2022["HCL"] = 1_424

	val := reflect.ValueOf(usage2022)
	assert.Equal(t, reflect.Map, val.Kind())

	keys := val.MapKeys()
	assert.Equal(t, reflect.String, keys[0].Kind())

	value := val.MapIndex(keys[0])
	assert.Equal(t, reflect.Int, value.Kind())

	for _, key := range val.MapKeys() {
		value = val.MapIndex(key)
		fmt.Printf("%s = %d\n", key.Interface(), value.Interface())
	}
}
