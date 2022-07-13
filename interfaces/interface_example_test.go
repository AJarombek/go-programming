/**
 * Test file for subtleties when satisfying interfaces
 * Author: Andrew Jarombek
 * Date: 7/10/2022
 */

package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

type Product interface {
	fmt.Stringer
	Purchase(transaction Transaction, count int) Product
	UniqueRequest(request string) Product
	Price() int
	Description() string
}

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

type Transaction interface {
	fmt.Stringer
	Complete() (ok bool)
	Add(product Product) Transaction
	AddAll(products ...Product) Transaction
	PaymentMethod(method PaymentMethod) Transaction
}

// ChristmasProduct conforms to the Product interface
type ChristmasProduct struct {
	price         int
	description   string
	uniqueRequest string
}

// ChristmasProduct methods

func (c ChristmasProduct) String() string {
	return fmt.Sprintf("%T for %d dollars.", c, c.price)
}

func (c ChristmasProduct) Purchase(transaction Transaction, count int) Product {
	for i := 0; i < count; i++ {
		transaction.Add(c)
	}

	return c
}

func (c ChristmasProduct) UniqueRequest(request string) Product {
	c.uniqueRequest = request
	return c
}

func (c ChristmasProduct) Price() int {
	return c.price
}

func (c ChristmasProduct) Description() string {
	return c.description
}

// ChristmasTree conforms to the Tree interface
type ChristmasTree struct {
	ChristmasProduct
	freshCut     bool
	trimBranches int
}

// ChristmasTree methods

func (c ChristmasTree) Height(height TreeHeight) Tree {
	c.price = int(height)
	return c
}

func (c ChristmasTree) GiveFreshCut(cut bool) Tree {
	c.freshCut = cut
	return c
}

func (c ChristmasTree) TrimBranches(count int) Tree {
	c.trimBranches = count
	return c
}

// HolidayWreath conforms to the Wreath interface
type HolidayWreath struct {
	ChristmasProduct
}

// HolidayWreath methods

func (wreath HolidayWreath) Size(size WreathSize) Wreath {
	wreath.price = int(size)
	return wreath
}

// HolidayRoping conforms to the Roping interface
type HolidayRoping struct {
	ChristmasProduct
}

// HolidayRoping methods

func (roping HolidayRoping) Length(length RopingLength) Roping {
	roping.price = int(length)
	return roping
}

// BundleGreens conforms to the Bundles interface
type BundleGreens struct {
	ChristmasProduct
}

// Greens methods

func (g BundleGreens) Count(size BundleSize, count int) Bundles {
	g.price = int(size) * count
	return g
}

// TreeLotPurchase conforms to the Transaction interface
type TreeLotPurchase struct {
	Price     int
	LineItems []Product
	Method    PaymentMethod
}

// TreeLotPurchase methods

func (purchase *TreeLotPurchase) String() string {
	return fmt.Sprintf("%d items for %d dollars", len(purchase.LineItems), purchase.Price)
}

func (purchase *TreeLotPurchase) Complete() (ok bool) {
	return true
}

func (purchase *TreeLotPurchase) Add(product Product) Transaction {
	purchase.Price += product.Price()
	purchase.LineItems = append(purchase.LineItems, product)
	return purchase
}

func (purchase *TreeLotPurchase) AddAll(products ...Product) Transaction {
	for _, product := range products {
		purchase.Price += product.Price()
		purchase.LineItems = append(purchase.LineItems, product)
	}

	return purchase
}

func (purchase *TreeLotPurchase) PaymentMethod(method PaymentMethod) Transaction {
	purchase.Method = method
	return purchase
}

func TestInterfaceExample(t *testing.T) {
	transaction := TreeLotPurchase{}

	// Basic use of the API to buy a tree
	transaction.Add(
		ChristmasTree{}.
			Height(Frasier6to7).
			GiveFreshCut(true).
			TrimBranches(0),
	).
		PaymentMethod(Cash).
		Complete()

	assert.Equal(t, 1, len(transaction.LineItems))
	assert.Equal(t, 70, transaction.Price)
	assert.Equal(t, Cash, transaction.Method)

	transaction = TreeLotPurchase{}
	tree := ChristmasTree{}.
		Height(Balsam9to10).
		GiveFreshCut(true).
		TrimBranches(2).
		UniqueRequest("Screw into a tree stand")

	doorWreath := HolidayWreath{}.Size(Decorated12Inch)
	mailboxWreath := HolidayWreath{}.Size(Plain6Inch)
	porchRoping := HolidayRoping{}.Length(Yards25)
	chimneyBranches := BundleGreens{}.Count(Standard, 2)

	transaction.
		AddAll(tree, doorWreath, mailboxWreath, porchRoping, chimneyBranches).
		PaymentMethod(Bitcoin).
		Complete()

	assert.Equal(t, Balsam9to10, tree.Price())
	assert.Equal(t, "interfaces.ChristmasProduct for 80 dollars.", tree.String())
	assert.Equal(t, "", tree.Description())

	assert.Equal(t, Decorated12Inch, doorWreath.Price())
	assert.Equal(t, "interfaces.ChristmasProduct for 17 dollars.", doorWreath.String())

	assert.Equal(t, Plain6Inch, mailboxWreath.Price())
	assert.Equal(t, "interfaces.ChristmasProduct for 8 dollars.", mailboxWreath.String())

	assert.Equal(t, Yards25, porchRoping.Price())
	assert.Equal(t, "interfaces.ChristmasProduct for 45 dollars.", porchRoping.String())

	assert.Equal(t, int(Standard)*2, chimneyBranches.Price())
	assert.Equal(t, "interfaces.ChristmasProduct for 6 dollars.", chimneyBranches.String())

	assert.Equal(t, 5, len(transaction.LineItems))
	assert.Equal(t, 156, transaction.Price)
	assert.Equal(t, PaymentMethod(Bitcoin), transaction.Method)
}
