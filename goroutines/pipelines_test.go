/**
 * Test file for basics of goroutine pipelines in Go
 * Author: Andrew Jarombek
 * Date: 7/16/2022
 */

package goroutines

import (
	"encoding/json"
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

type Exercises []Exercise

func TestPipelines(t *testing.T) {
	extract := make(chan string)
	raw := make(chan *string)
	mapp := make(chan Exercises)
	filter := make(chan Exercises)
	reduce := make(chan float64)

	// Extract data
	go func() {
		filename := <-extract
		data, err := os.ReadFile(filename)

		if err != nil {
			assert.Fail(t, err.Error())
		}

		dataString := string(data)
		raw <- &dataString
	}()

	// Map the raw data
	go func() {
		var exercises Exercises
		data := <-raw

		err := json.Unmarshal([]byte(*data), &exercises)

		if err != nil {
			assert.Fail(t, err.Error())
		}

		mapp <- exercises
	}()

	// Filter the mapped data
	go func() {
		exercises := <-mapp

		index := 0
		end := 0

		for index < len(exercises) {
			if exercises[index].Type == "run" {
				exercises[end] = exercises[index]
				end++
			}

			index++
		}

		filter <- exercises[:end]
	}()

	// Reduce the mapped data
	go func() {
		exercises := <-filter
		result := 0.0

		for _, exercise := range exercises {
			result += exercise.Miles
		}

		reduce <- result
	}()

	extract <- "./exercises.json"
	result := <-reduce
	assert.Equal(t, 6.76, result)
}
