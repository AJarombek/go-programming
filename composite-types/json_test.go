/**
 * Test file for creating JSON from structs and converting structs to JSON
 * Author: Andrew Jarombek
 * Date: 7/1/2022
 */

package _go

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ConferenceSession struct {
	Title        string
	Description  string
	Attended     bool     `json:"attended,omitempty"`
	Technologies []string `json:"technology_list"`
}

func TestJSON(t *testing.T) {
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
		assert.Fail(t, "Failed to convert struct to JSON")
	}

	assert.Equal(
		t,
		"{\"Title\":\"Migrate Your Existing DAGs to Databricks Workflows\",\"Description\":\"\",\"attended\":true,\"technology_list\":[\"Databricks\",\"Airflow\"]}",
		string(data),
	)

	// When attended = false, it is omitted from the JSON string
	data, err = json.Marshal(sessions[1])

	if err != nil {
		assert.Fail(t, "Failed to convert struct to JSON")
	}

	assert.Equal(
		t,
		"{\"Title\":\"The Road to a Robust Data Lake: Utilizing Delta Lake and Databricks to Map 150 Million Miles of Roads a Month\",\"Description\":\"\",\"technology_list\":[\"Databricks\"]}",
		string(data),
	)

	data, err = json.Marshal(sessions)

	if err != nil {
		assert.Fail(t, "Failed to convert list of structs to JSON")
	}

	var newSessions []ConferenceSession

	if err = json.Unmarshal(data, &newSessions); err != nil {
		assert.Fail(t, "Failed to convert JSON to a struct")
	}

	assert.Equal(t, 3, len(newSessions))
	assert.Equal(t, "Talk about my work at Cigna", newSessions[2].Description)

	var sessionTitles []struct{ Title string }

	if err = json.Unmarshal(data, &sessionTitles); err != nil {
		assert.Fail(t, "Failed to convert JSON to a struct")
	}

	assert.Equal(t, 3, len(sessionTitles))
	assert.Equal(t, struct{ Title string }{
		Title: "Journey to Solving Healthcare Price Transparency with Databricks and Delta Lake",
	}, sessionTitles[2])
}
