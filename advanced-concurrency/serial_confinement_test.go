/**
 * Test file for using the serial confinement pattern in concurrent code
 * Author: Andrew Jarombek
 * Date: 7/31/2022
 */

package advanced_concurrency

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"testing"
)

// Serial confinement is when channel pipelines are used in concurrent code,
// and variables are confined in a single stage of the pipeline at a time
// before being passed to the next stage.  The example below shows the creation
// of a blanket.
type YarnWeight int

const (
	SuperFine YarnWeight = iota + 1
	Fine
	Light
	Medium
	Bulky
	SuperBulky
	Jumbo
)

type Blanket struct {
	colors         []color.Color
	yarnWeight     YarnWeight
	knitted        bool
	weaveLooseEnds bool
	givenLove      bool
}

var newBlanket = make(chan struct{})

func design(designed chan<- *Blanket) {
	for range newBlanket {
		blanket := new(Blanket)

		// This can be made more flexible, currently it will always create a
		// blanket of the same design, like a factory.
		// Color palette https://colorhunt.co/palette/377d71fbc5c5fba1a18879b0
		blanket.yarnWeight = 5
		blanket.colors = []color.Color{
			color.RGBA{R: 55, G: 125, B: 113},
			color.RGBA{R: 251, G: 197, B: 197},
			color.RGBA{R: 251, G: 161, B: 161},
			color.RGBA{R: 136, G: 121, B: 176},
		}

		designed <- blanket
	}
}

func knit(finished chan<- *Blanket, designed <-chan *Blanket) {
	for blanket := range designed {
		blanket.knitted = true
		blanket.weaveLooseEnds = true
		blanket.givenLove = true
		finished <- blanket
	}
}

func TestSerialConfinement(t *testing.T) {
	designed := make(chan *Blanket)
	finished := make(chan *Blanket)

	go design(designed)
	go knit(finished, designed)

	assert.Equal(t, 0, len(finished))

	newBlanket <- struct{}{}

	blankie := <-finished
	assert.True(t, blankie.knitted)
	assert.True(t, blankie.weaveLooseEnds)
	assert.True(t, blankie.givenLove)

	assert.Equal(t, YarnWeight(5), blankie.yarnWeight)
}
