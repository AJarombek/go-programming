/**
 * Test file for creating parameterized text files via a template
 * Author: Andrew Jarombek
 * Date: 7/2/2022
 */

package _go

import (
	"github.com/stretchr/testify/assert"
	"html/template"
	"os"
	"testing"
)

const templ = `
{{range .}}
{{.Name}} - {{.Type}} on {{.Date}}
{{.Miles}}mi
{{.Time}}
{{- if not .Description -}}
{{- else }}
{{.Description -}}
{{ end }}
{{end}}
`

type Log struct {
	Name        string
	Type        string
	Date        string
	Miles       float64
	Time        string
	Description string
}

var report = template.Must(template.New("logs").Parse(templ))

func TestTemplate(t *testing.T) {
	logs := []Log{
		{Name: "To Grand Central", Type: "Walk", Date: "July 2nd, 2022", Miles: 1.73, Time: "33:31", Description: ""},
		{Name: "Central Park Heat", Type: "Run", Date: "July 2nd, 2022", Miles: 6.10, Time: "44:49",
			Description: "Body not used to the heat anymore"},
		{Name: "Pre-Flight", Type: "Run", Date: "July 1st, 2022", Miles: 4.42, Time: "35:31"},
		{Name: "Presidio Hills", Type: "Bike", Date: "June 30th, 2022", Miles: 13, Time: "1:21:46",
			Description: "Beautiful"},
	}

	if err := report.Execute(os.Stdout, logs); err != nil {
		assert.Fail(t, err.Error())
	}
}
