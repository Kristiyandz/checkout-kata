package checkout

import (
	"fmt"
	"log"

	"github.com/Kristiyandz/checkout-kata/model"
)

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type CheckoutBasket struct {
	items        []model.Item
	pricingRules model.PricingRules
}

// initi a new checout with provided pricing rules
func InitNewCheckout(pricingRules model.PricingRules) ICheckout {
	return &CheckoutBasket{
		pricingRules: pricingRules,
	}
}

// scan adds items to the checkout basked
func (c *CheckoutBasket) Scan(itemSKU string) {

	unitPrice, exists := c.pricingRules.Prices[itemSKU]

	if !exists {
		log.Fatal("Item SKU not found in pricing rules")
	}

	basketitem := model.Item{SKU: itemSKU, UnitPrice: unitPrice}
	c.items = append(c.items, basketitem)
}

// gets the total price
func (c *CheckoutBasket) GetTotalPrice() int {
	fmt.Println("Getting total value")
	return 0
}
