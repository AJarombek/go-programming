/**
 * Test file for using sort.Interface
 * Author: Andrew Jarombek
 * Date: 7/12/2022
 */

package interfaces

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type Walk Run

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

func TestSortingInterface(t *testing.T) {
	walks := []*Walk{
		{Miles: 1.79, Name: "Central Park Ramble"},
		{Miles: 3.34, Name: "Central Park + Deli Stop"},
		{Miles: 2.06, Name: "Central Park South"},
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
}
