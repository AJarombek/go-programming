/**
 * Main test file for composite types
 * Author: Andrew Jarombek
 * Date: 6/27/2022
 */

package _go

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Exercise struct {
	Hours, Minutes, Seconds uint8
	Miles                   float64
	Feel                    uint8
	Name                    string
	Description             string
}

type Run struct {
	Exercise
	Shoe string
}

type Bike struct {
	Exercise
	Model string
}

type ConferenceSession struct {
	Title        string
	Description  string
	Attended     bool     `json:"attended,omitempty"`
	Technologies []string `json:"technology_list"`
}

func Test(t *testing.T) {
	// Struct embedding allows nested fields to be assigned to and accessed
	// on the top level struct variable.
	var easyRun Run
	easyRun.Hours = 0
	easyRun.Minutes = 56
	easyRun.Seconds = 50
	easyRun.Miles = 7.66
	easyRun.Feel = 6
	easyRun.Name = "Cognewaugh & Cat Rock"
	easyRun.Description = "Tired from lots of driving, saw some deer which made the run worthwhile"
	easyRun.Shoe = "Asics GT-2000"

	assert.Equal(t, uint8(56), easyRun.Minutes)
	assert.Equal(t, easyRun.Exercise.Name, easyRun.Name)

	// Another way to create an embedded struct
	hillyRun := Run{
		Exercise: Exercise{
			Hours:       0,
			Minutes:     46,
			Seconds:     33,
			Miles:       5.64,
			Feel:        4,
			Name:        "Exploring SF",
			Description: "Late flight so only got 4 hours of sleep, very tired.  Tech conference starts tomorrow",
		},
		Shoe: "Asics GT-2000",
	}

	assert.Equal(t, "Exploring SF", hillyRun.Name)

	sessions := []ConferenceSession{
		{Title: "Migrate Your Existing DAGs to Databricks Workflows", Attended: true,
			Technologies: []string{"Databricks", "Airflow"}},
		{Title: "The Road to a Robust Data Lake: Utilizing Delta Lake and Databricks to Map 150 Million Miles of Roads a Month",
			Attended: false, Technologies: []string{"Databricks"}},
		{Title: "Journey to Solving Healthcare Price Transparency with Databricks and Delta Lake",
			Description: "Talk about my work at Cigna", Attended: true, Technologies: []string{"Databricks", "Airflow"}},
	}

	data, err := json.Marshal(sessions[0])

	if err != nil {
		log.Fatalf("Failed to convert struct to JSON: %s", err)
	}

	assert.Equal(
		t,
		"\"Title\":\"Migrate Your Existing DAGs to Databricks Workflows\",\"Description\":\"\",\"attended\":true,\"technology_list\":[\"Databricks\",\"Airflow\"]}",
		data,
	)
}
