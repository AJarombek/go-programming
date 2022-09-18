/**
 * Test file for unidirectional channels in Go
 * Author: Andrew Jarombek
 * Date: 7/17/2022
 */

package goroutines

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func read(out chan<- Exercises) {
	filename := "./exercises.json"

	if os.Getenv("TEST_ENV") == "plz" {
		filename = "./goroutines/exercises.json"
	}

	data, err := os.ReadFile(filename)

	if err != nil {
		out <- Exercises{}
	}

	var exercises Exercises
	err = json.Unmarshal(data, &exercises)

	if err != nil {
		out <- Exercises{}
	}

	out <- exercises
}

func filter(out chan<- Exercises, in <-chan Exercises) {
	exercises := <-in

	index := 0
	end := 0

	for index < len(exercises) {
		if exercises[index].Type == "kayak" && exercises[index].Miles > 5 {
			exercises[end] = exercises[index]
			end++
		}

		index++
	}

	out <- exercises[:end]
}

func reduce(out chan<- float64, in <-chan Exercises) {
	exercises := <-in
	result := 0.0

	for _, exercise := range exercises {
		result += exercise.Miles
	}

	out <- result
}

func TestUnidirectionalChannels(t *testing.T) {
	readChan := make(chan Exercises)
	filterChan := make(chan Exercises)
	reduceChan := make(chan float64)

	go read(readChan)
	go filter(filterChan, readChan)
	go reduce(reduceChan, filterChan)

	assert.Equal(t, 11.04, <-reduceChan)

	close(readChan)
	close(filterChan)
	close(reduceChan)
}
