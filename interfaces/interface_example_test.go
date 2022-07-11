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

type TreeHeight int

const (
	Balsam5to6    TreeHeight = 40
	Balsam7to8               = 50
	Balsam8to9               = 65
	Balsam9to10              = 80
	Balsam10to12             = 120
	Frasier4to5              = 35
	Frasier5to6              = 60
	Frasier6to7              = 70
	Frasier7to8              = 80
	Frasier8to9              = 95
	Frasier9to10             = 120
	Frasier10to12            = 150
	Douglas5to6              = 55
	Douglas6to7              = 65
	Douglas7to8              = 75
	Douglas8to9              = 90
)

type WreathSize int

const (
	DecoratedSpray  WreathSize = 10
	Decorated6Inch             = 12
	Decorated8Inch             = 14
	Decorated12Inch            = 17
	Decorated15Inch            = 24
	Decorated18Inch            = 35
	Decorated24Inch            = 45
	Plain6Inch                 = 8
	Plain8Inch                 = 10
	Plain12Inch                = 13
	Plain15Inch                = 17
	Plain18Inch                = 25
	Plain24Inch                = 35
	Plain30Inch                = 45
	Plain36Inch                = 55
)

type RopingLength int

const (
	Feet25  RopingLength = 20
	Yards25              = 45
)

type BundleSize int

const (
	Standard BundleSize = 3
)

type PaymentMethod string

const (
	Cash    PaymentMethod = "cash"
	Check                 = "check"
	Bitcoin               = "bitcoin" // USB Drives Only
)

type Bundles interface {
	Product
	Count(size BundleSize, count int) Bundles
}

type Roping interface {
	Product
	Length(length RopingLength) Roping
}

type Wreath interface {
	Product
	Size(size WreathSize) Wreath
}

type Tree interface {
	Product
	Height(height TreeHeight) Tree
	GiveFreshCut(cut bool) Tree
	TrimBranches(count int) Tree
}

type Product interface {
	fmt.Stringer
	Purchase(transaction Transaction, count int) Product
	UniqueRequest(request string) Product
	Price() int
	Description() string
}

type Transaction interface {
	fmt.Stringer
	Complete() (ok bool)
	Add(product Product) Transaction
	AddAll(products ...Product) Transaction
	PaymentMethod(method PaymentMethod) Transaction
}

type TreeLotPurchase struct {
	Price     int
	LineItems []Product
	Method    PaymentMethod
}

func (purchase TreeLotPurchase) String() string {
	return fmt.Sprintf("%d items for %d dollars", len(purchase.LineItems), purchase.Price)
}

func (purchase TreeLotPurchase) Complete() (ok bool) {
	return true
}

func (purchase TreeLotPurchase) Add(product Product) Transaction {
	purchase.Price += product.Price()
	purchase.LineItems = append(purchase.LineItems, product)
	return purchase
}

func (purchase TreeLotPurchase) AddAll(products ...Product) Transaction {
	for _, product := range products {
		purchase.Price += product.Price()
		purchase.LineItems = append(purchase.LineItems, product)
	}

	return purchase
}

func (purchase TreeLotPurchase) PaymentMethod(method PaymentMethod) Transaction {
	purchase.Method = method
	return purchase
}

func TestInterfaceExample(t *testing.T) {
	// transaction := TreeLotPurchase{}
}
