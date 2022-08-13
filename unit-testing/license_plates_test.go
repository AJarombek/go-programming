/**
 * Test file for the license_plates.go file
 * Author: Andrew Jarombek
 * Date: 8/8/2022
 */

package unit_testing

import "testing"

func TestValidCountryCode(t *testing.T) {
	code, ok := CountryCode("Slovakia")

	if code != "SK" {
		t.Errorf(`countryCode("Slovakia") != "SK"`)
	}

	if !ok {
		t.Errorf(`countryCode("Slovakia") not ok`)
	}
}

func TestInvalidCountryCode(t *testing.T) {
	// European country license plate dictionary excludes intercontinental countries
	code, ok := CountryCode("Turkey")

	if code != "" {
		t.Errorf(`CountryCode("Turkey") != ""`)
	}

	if ok {
		t.Errorf(`CountryCode("Turkey") ok`)
	}
}

func TestCollected(t *testing.T) {
	// Since t.Errorf() does not crash immediately after the first failure,
	// it is okay to use it in a loop over test conditions.
	tests := []struct {
		input    string
		expected int
		ok       bool
	}{
		{"United Kingdom", 0, true},
		{"Ukraine", 5, true},
		{"Slovakia", 2, true},
		{"Kazakhstan", 0, false},
		{"France", 1, true},
	}

	for _, test := range tests {
		collected, ok := Collected(test.input)

		if collected != test.expected {
			t.Errorf("Collected(%q) = %d, want %d", test.input, collected, test.expected)
		}

		if ok != test.ok {
			t.Errorf("countryCode(%q) ok = %v, want %v", test.input, ok, test.ok)
		}
	}
}

func BenchmarkCountryCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountryCode("Slovakia")
	}
}

func benchmarkCollected(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			Collected("Ukraine")
		}
	}
}

func BenchmarkCollected10(b *testing.B) {
	benchmarkCollected(b, 10)
}

func BenchmarkCollected100(b *testing.B) {
	benchmarkCollected(b, 100)
}

func BenchmarkCollected1000(b *testing.B) {
	benchmarkCollected(b, 1000)
}
