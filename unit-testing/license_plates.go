/**
 * Functions for information about license plates.
 * Author: Andrew Jarombek
 * Date: 8/8/2022
 */

package unit_testing

var europeLpCodes = map[string]string{
	"Albania":                "AL",
	"Andorra":                "AND",
	"Austria":                "A",
	"Belarus":                "BY",
	"Belgium":                "B",
	"Bosnia and Herzegovina": "BIH",
	"Bulgaria":               "BG",
	"Croatia":                "HR",
	"Cyprus":                 "CY",
	"Czech Republic":         "CZ",
	"Denmark":                "DK",
	"Estonia":                "EST",
	"Finland":                "FIN",
	"France":                 "F",
	"Germany":                "D",
	"Greece":                 "GR",
	"Hungary":                "H",
	"Iceland":                "IS",
	"Ireland":                "IRL",
	"Italy":                  "I",
	"Latvia":                 "LV",
	"Liechtenstein":          "FL",
	"Lithuania":              "LT",
	"Luxembourg":             "L",
	"Malta":                  "M",
	"Moldova":                "MD",
	"Monaco":                 "MC",
	"Montenegro":             "MNE",
	"Netherlands":            "NL",
	"North Macedonia":        "NMK",
	"Norway":                 "N",
	"Poland":                 "PL",
	"Portugal":               "P",
	"Romania":                "RO",
	"San Marino":             "",
	"Serbia":                 "SRB",
	"Slovakia":               "SK",
	"Slovenia":               "SLO",
	"Spain":                  "E",
	"Sweden":                 "S",
	"Switzerland":            "",
	"Ukraine":                "UA",
	"United Kingdom":         "UK",
	"Vatican City":           "",
}

func countryCode(country string) (string, bool) {
	if code, ok := europeLpCodes[country]; ok {
		return code, true
	}

	return "", false
}
