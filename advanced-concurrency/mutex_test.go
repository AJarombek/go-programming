/**
 * Test file for using binary semaphores and sync.Mutex in concurrent code
 * Author: Andrew Jarombek
 * Date: 7/31/2022
 */

package advanced_concurrency

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var (
	employee  = Employee{sickDays: 10}
	semaphore = make(chan struct{}, 1)
)

// A binary semaphore can be used to ensure that at most
// one goroutine can access a shared variable at a time.

func CheckBalanceV2() int {
	semaphore <- struct{}{}
	balance := employee.sickDays
	<-semaphore
	return balance
}

func UseSickDaysV2(count int) {
	semaphore <- struct{}{}
	employee.sickDays -= count
	<-semaphore
}

// The binary semaphore pattern is the same as the mutual exclusion pattern,
// which is implemented in Go with the type sync.Mutex.  The following two methods
// are equivalent to the V2 methods above.

var mutex sync.Mutex

func CheckBalanceV3() int {
	mutex.Lock()
	balance := employee.sickDays
	mutex.Unlock()
	return balance
}

func UseSickDaysV3(count int) {
	mutex.Lock()
	employee.sickDays -= count
	mutex.Unlock()
}

// These functions can be improved slightly by deferring
// the call to Unlock() on the mutex.

func CheckBalanceV4() int {
	mutex.Lock()
	defer mutex.Unlock()
	return employee.sickDays
}

func UseSickDaysV4(count int) {
	mutex.Lock()
	defer mutex.Unlock()
	employee.sickDays -= count
}

func TestMutex(t *testing.T) {
	balance := CheckBalanceV2()
	assert.Equal(t, 10, balance)

	UseSickDaysV2(1)
	balance = CheckBalanceV2()
	assert.Equal(t, 9, balance)

	UseSickDaysV3(1)
	balance = CheckBalanceV3()
	assert.Equal(t, 8, balance)

	UseSickDaysV4(1)
	balance = CheckBalanceV4()
	assert.Equal(t, 7, balance)
}
