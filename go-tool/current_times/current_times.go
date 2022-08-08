/**
 * Basic Go code to print the current time in different locations.
 * Author: Andrew Jarombek
 * Date: 8/5/2022
 */

package main

import (
	"fmt"
	"time"
)

// currentTime returns a string representation of the current time in a time zone.
func currentTime(timeZone string) (string, error) {
	location, err := time.LoadLocation(timeZone)

	if err != nil {
		return "", fmt.Errorf("invalid time zone")
	}

	return time.Now().In(location).Format("3:04 PM Jan 2 2006"), nil
}

func main() {
	zones := []string{
		"America/New_York",
		"America/Denver",
		"Europe/London",
		"Europe/Istanbul",
	}

	locations := []string{
		"New York",
		"Denver",
		"London",
		"Istanbul",
	}

	for index, zone := range zones {
		ct, err := currentTime(zone)

		if err == nil {
			fmt.Println(locations[index] + " " + ct)
		}
	}
}
