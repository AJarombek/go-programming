/**
 * Test file for basics of goroutine pipelines in Go
 * Author: Andrew Jarombek
 * Date: 7/16/2022
 */

package goroutines

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Exercise struct {
	Name        string  `json:"name"`
	Date        string  `json:"date"`
	Type        string  `json:"type"`
	Miles       float64 `json:"miles,omitempty"`
	Time        string  `json:"time,omitempty"`
	Description string  `json:"description,omitempty"`
}

func TestPipelines(t *testing.T) {
	extract := make(chan string)
	raw := make(chan *string)
	mapp := make(chan *Exercise)
	filter := make(chan *Exercise)
	reduce := make(chan float64)

	// Extract data
	go func() {
		filename := <-extract
		data, err := os.ReadFile(filename)

		if err != nil {
			assert.Fail(t, err.Error())
		}
		
		extract <- string(data)
	}()

	// Map the raw data
	go func() {}()

	// Filter the mapped data
	go func() {}()

	// Reduce the mapped data
	go func() {}()

	extract <- "./exercises.json"
	result := <-reduce
}