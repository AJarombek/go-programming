/**
 * Test file for using the monitor goroutine pattern in concurrent code
 * Author: Andrew Jarombek
 * Date: 7/31/2022
 */

package advanced_concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee struct {
	sickDays int
}

var balance = make(chan int)
var use = make(chan int)

func CheckBalance() int {
	return <-balance
}

func UseSickDays(count int) {
	use <- count
}

// A monitor goroutine is a goroutine that handles read and write access to a confined variable via channels.
// This allows the confined variable to be used in concurrent programming without the risk of race conditions.
// This example is an HR platform handling sick days for an employee, inspired by my now week-long illness.
func monitor() {
	employee := Employee{sickDays: 10}

	for {
		select {
		case days := <-use:
			employee.sickDays -= days
		case balance <- employee.sickDays:
		}
	}
}

func TestMonitorGoroutine(t *testing.T) {
	go monitor()

	balance := CheckBalance()
	assert.Equal(t, 10, balance)

	UseSickDays(2)

	balance = CheckBalance()
	assert.Equal(t, 8, balance)
}
