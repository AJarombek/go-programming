/**
 * Test file for subtleties when satifying interfaces
 * Author: Andrew Jarombek
 * Date: 7/10/2022
 */

package interfaces

import (
	"fmt"
	"testing"
)

func (run *Run) Format(state fmt.State, verb rune) {
	switch verb {
	case 's':
		_, err := fmt.Fprint(state, run.String())

		if err != nil {
			return
		}
	case 'v':
		_, err := fmt.Fprintf(
			state,
			"{Miles=%.02f, Time=%d:%02d:%02d, Feel=%d, Name=%s, Description=%s}",
			run.Miles,
			run.Hours,
			run.Minutes,
			run.Seconds,
			run.Feel,
			run.Name,
			run.Description,
		)
		
		if err != nil {
			return
		}
	}
}

func TestInterfaceSatisfaction(t *testing.T) {

}
