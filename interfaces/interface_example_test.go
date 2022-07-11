/**
 * Test file for subtleties when satisfying interfaces
 * Author: Andrew Jarombek
 * Date: 7/10/2022
 */

package interfaces

import (
	"fmt"
	"testing"
)

type Bundle interface {
	Product
}

type Roping interface {
	Product
}

type Wreath interface {
	Product
}

type Tree interface {
	Product
}

type Product interface {
	fmt.Stringer
	Purchase(transaction Transaction)
	Price() float64
	Description() string
}

type Transaction interface {
	fmt.Stringer
	Create()
}

func TestInterfaceExample(t *testing.T) {

}
