/**
 * Test file for creating parameterized HTML files via a template
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

const htmlTempl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>
<body>
    <div>
	<h1>Mileage for Next {{.WeekCount}} Weeks</h1>
	<table>
	<tr>
		<th>Week</th>
		<th>Mileage</th>
	</tr>
	{{- range .Mileage}}
	<tr>
		<td>{{.Week}}</td>
		<td>{{.Miles}}</td>
	</tr>
	{{- end }}
	</table>
	</div>
</body>
</html>
`

type WeeklyMileage struct {
	Week  string
	Miles float64
}

type MileageReport struct {
	Mileage   []WeeklyMileage
	WeekCount int
}

var htmlReport = template.Must(template.New("logs").Parse(htmlTempl))

func TestHtmlTemplate(t *testing.T) {
	mileage := []WeeklyMileage{
		{Week: "7/4/2022 - 7/10/2022", Miles: 10},
		{Week: "7/11/2022 - 7/17/2022", Miles: 30},
		{Week: "7/18/2022 - 7/24/2022", Miles: 46},
		{Week: "7/25/2022 - 7/31/2022", Miles: 50},
	}
	mileageReport := MileageReport{
		Mileage:   mileage,
		WeekCount: len(mileage),
	}

	if err := htmlReport.Execute(os.Stdout, mileageReport); err != nil {
		assert.Fail(t, err.Error())
	}

	buf := &bytes.Buffer{}

	// Write the report to a buffer, which can be read in code
	if err := htmlReport.Execute(buf, mileageReport); err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>
<body>
    <div>
	<h1>Mileage for Next 4 Weeks</h1>
	<table>
	<tr>
		<th>Week</th>
		<th>Mileage</th>
	</tr>
	<tr>
		<td>7/4/2022 - 7/10/2022</td>
		<td>10</td>
	</tr>
	<tr>
		<td>7/11/2022 - 7/17/2022</td>
		<td>30</td>
	</tr>
	<tr>
		<td>7/18/2022 - 7/24/2022</td>
		<td>46</td>
	</tr>
	<tr>
		<td>7/25/2022 - 7/31/2022</td>
		<td>50</td>
	</tr>
	</table>
	</div>
</body>
</html>
`, buf.String())
}
