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

func main() {
	newYork, _ := time.LoadLocation("America/New_York")
	denver, _ := time.LoadLocation("America/Denver")
	london, _ := time.LoadLocation("Europe/London")
	istanbul, _ := time.LoadLocation("Europe/Istanbul")

	fmt.Println("New York " + time.Now().In(newYork).Format("3:04 PM Jan 2 2006"))
	fmt.Println("Denver   " + time.Now().In(denver).Format("3:04 PM Jan 2 2006"))
	fmt.Println("London   " + time.Now().In(london).Format("3:04 PM Jan 2 2006"))
	fmt.Println("Istanbul " + time.Now().In(istanbul).Format("3:04 PM Jan 2 2006"))
}
