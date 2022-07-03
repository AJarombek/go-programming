/**
 * Test file for creating parameterized text via a template
 * Author: Andrew Jarombek
 * Date: 7/2/2022
 */

package _go

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"html/template"
	"os"
	"testing"
)

const templ = `
{{- range .}}
{{.Name}} - {{.Type}} on {{.Date}}
{{.Miles}}mi
{{.Time}}
{{- if not .Description -}}
{{- else }}
{{.Description -}}
{{ end }}
{{end -}}
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

	// Write the report to standard output
	if err := report.Execute(os.Stdout, logs); err != nil {
		assert.Fail(t, err.Error())
	}

	buf := &bytes.Buffer{}

	// Write the report to a buffer, which can be read in code
	if err := report.Execute(buf, logs); err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, `
To Grand Central - Walk on July 2nd, 2022
1.73mi
33:31

Central Park Heat - Run on July 2nd, 2022
6.1mi
44:49
Body not used to the heat anymore

Pre-Flight - Run on July 1st, 2022
4.42mi
35:31

Presidio Hills - Bike on June 30th, 2022
13mi
1:21:46
Beautiful
`, buf.String())
}
