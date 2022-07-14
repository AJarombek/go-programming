/**
 * Test file for using sort.Interface
 * Author: Andrew Jarombek
 * Date: 7/12/2022
 */

package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type Walk Run

func (walk Walk) String() string {
	return fmt.Sprintf(
		"%s - %.02f miles walked in %d:%02d:%02d",
		walk.Name,
		walk.Miles,
		walk.Hours,
		walk.Minutes,
		walk.Seconds,
	)
}

type byMiles []*Walk

func (x byMiles) Len() int {
	return len(x)
}

func (x byMiles) Less(i, j int) bool {
	return x[j].Miles < x[i].Miles
}

func (x byMiles) Swap(i, j int) {
	x[j], x[i] = x[i], x[j]
}

type byTime []*Walk

func (x byTime) Len() int {
	return len(x)
}

func (x byTime) Less(i, j int) bool {
	if x[j].Hours != x[i].Hours {
		return x[i].Hours < x[j].Hours
	}

	if x[j].Minutes != x[i].Minutes {
		return x[i].Minutes < x[j].Minutes
	}

	return x[i].Seconds < x[j].Seconds
}

func (x byTime) Swap(i, j int) {
	x[j], x[i] = x[i], x[j]
}

func TestSortingInterface(t *testing.T) {
	walks := []*Walk{
		{Miles: 1.79, Hours: 0, Minutes: 35, Seconds: 15, Name: "Central Park Ramble"},
		{Miles: 3.34, Hours: 1, Minutes: 2, Seconds: 51, Name: "Central Park + Deli Stop"},
		{Miles: 2.06, Hours: 0, Minutes: 35, Seconds: 10, Name: "Central Park South"},
	}

	assert.Equal(t, 1.79, walks[0].Miles)
	assert.Equal(t, 3.34, walks[1].Miles)
	assert.Equal(t, 2.06, walks[2].Miles)

	sort.Sort(byMiles(walks))

	assert.Equal(t, 3.34, walks[0].Miles)
	assert.Equal(t, 2.06, walks[1].Miles)
	assert.Equal(t, 1.79, walks[2].Miles)

	sort.Sort(sort.Reverse(byMiles(walks)))

	assert.Equal(t, 1.79, walks[0].Miles)
	assert.Equal(t, 2.06, walks[1].Miles)
	assert.Equal(t, 3.34, walks[2].Miles)

	sort.Sort(byTime(walks))

	assert.Equal(t, 2.06, walks[0].Miles)
	assert.Equal(t, 1.79, walks[1].Miles)
	assert.Equal(t, 3.34, walks[2].Miles)
}
