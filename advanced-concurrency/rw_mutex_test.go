/**
 * Test file for using sync.RWMutex in concurrent code
 * Author: Andrew Jarombek
 * Date: 8/2/2022
 */

package advanced_concurrency

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var (
	contractor = Employee{sickDays: 0}
	rwMutex    sync.RWMutex
)

// The RWMutex (read write mutex) differs from the standard Mutex because
// it is a lock that allows multiple concurrent readers, but only one concurrent
// writer.  There can not be a concurrent reader and writer.  The reader lock uses
// the RLock() method while the writer lock still uses the Lock() method.

func CheckBalanceV5() int {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return contractor.sickDays
}

func UseSickDaysV5(count int) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	contractor.sickDays -= count
}

// A more complex example uses both a read lock and a write lock within the same function.
// This example lazy loads a dictionary in a concurrency safe manner.
var mu sync.RWMutex

type SnuggleBug struct {
	name    string
	species string
}

var snuggleFriends map[string]SnuggleBug

func populateSnuggleFriends() {
	snuggleFriends = make(map[string]SnuggleBug)
	snuggleFriends["dotty"] = SnuggleBug{
		name:    "Dotty",
		species: "Horse",
	}
	snuggleFriends["lily"] = SnuggleBug{
		name:    "Lily",
		species: "Bear",
	}
}

func GetSnuggleFriend(name string) SnuggleBug {
	mu.RLock()
	if snuggleFriends != nil {
		friend := snuggleFriends[name]
		mu.RUnlock()
		return friend
	}

	mu.Lock()
	if snuggleFriends == nil {
		populateSnuggleFriends()
	}

	friend := snuggleFriends[name]
	mu.Unlock()

	return friend
}

// Equivalent to the above GetSnuggleFriend() function,
// except using sync.Once instead of sync.RWMutex.

var loadFriendsOnce sync.Once

func GetSnuggleFriendV2(name string) SnuggleBug {
	loadFriendsOnce.Do(populateSnuggleFriends)
	return snuggleFriends[name]
}

func TestRWMutex(t *testing.T) {
	balance := CheckBalanceV5()
	assert.Equal(t, 0, balance)

	UseSickDaysV5(1)
	balance = CheckBalanceV5()
	assert.Equal(t, -1, balance)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		assert.Equal(t, "Dotty", GetSnuggleFriend("dotty").name)
		wg.Done()
	}()

	go func() {
		assert.Equal(t, "Bear", GetSnuggleFriend("lily").species)
		wg.Done()
	}()

	go func() {
		assert.Equal(t, "Horse", GetSnuggleFriendV2("dotty").species)
		wg.Done()
	}()

	go func() {
		assert.Equal(t, "Lily", GetSnuggleFriendV2("lily").name)
		wg.Done()
	}()

	wg.Wait()
	assert.NotNil(t, snuggleFriends)
}
