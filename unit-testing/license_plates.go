/**
 * Functions for information about license plates.
 * Based off data from https://en.wikipedia.org/wiki/Vehicle_registration_plates_of_Europe
 * Author: Andrew Jarombek
 * Date: 8/8/2022
 */

package unit_testing

type EuropeLP struct {
	code      string
	collected int
}

var europeLpCodes = map[string]EuropeLP{
	"Albania":                {code: "AL"},
	"Andorra":                {code: "AND"},
	"Austria":                {code: "A", collected: 1},
	"Belarus":                {code: "BY"},
	"Belgium":                {code: "B"},
	"Bosnia and Herzegovina": {code: "BIH"},
	"Bulgaria":               {code: "BG"},
	"Croatia":                {code: "HR"},
	"Cyprus":                 {code: "CY"},
	"Czech Republic":         {code: "CZ", collected: 2},
	"Denmark":                {code: "DK"},
	"Estonia":                {code: "EST"},
	"Finland":                {code: "FIN"},
	"France":                 {code: "F", collected: 1},
	"Germany":                {code: "D", collected: 2},
	"Greece":                 {code: "GR"},
	"Hungary":                {code: "H"},
	"Iceland":                {code: "IS"},
	"Ireland":                {code: "IRL"},
	"Italy":                  {code: "I", collected: 1},
	"Latvia":                 {code: "LV"},
	"Liechtenstein":          {code: "FL"},
	"Lithuania":              {code: "LT"},
	"Luxembourg":             {code: "L"},
	"Malta":                  {code: "M"},
	"Moldova":                {code: "MD"},
	"Monaco":                 {code: "MC"},
	"Montenegro":             {code: "MNE"},
	"Netherlands":            {code: "NL", collected: 1},
	"North Macedonia":        {code: "NMK"},
	"Norway":                 {code: "N"},
	"Poland":                 {code: "PL", collected: 3},
	"Portugal":               {code: "P"},
	"Romania":                {code: "RO"},
	"San Marino":             {},
	"Serbia":                 {code: "SRB"},
	"Slovakia":               {code: "SK", collected: 2},
	"Slovenia":               {code: "SLO"},
	"Spain":                  {code: "E"},
	"Sweden":                 {code: "S"},
	"Switzerland":            {collected: 1},
	"Ukraine":                {code: "UA", collected: 5},
	"United Kingdom":         {code: "UK"},
	"Vatican City":           {},
}

// CountryCode get the country code displayed on a european license plate
func CountryCode(country string) (string, bool) {
	if lp, ok := europeLpCodes[country]; ok {
		return lp.code, true
	}

	return "", false
}

// Collected get the number of license plates I've collected from a certain country
func Collected(country string) (int, bool) {
	if lp, ok := europeLpCodes[country]; ok {
		return lp.collected, true
	}

	return 0, false
}
