package checkout

import (
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
	totalPrice := 0
	itemCounts := make(map[string]int)

	for _, item := range c.items {
		itemCounts[item.SKU]++
	}

	for sku, count := range itemCounts {
		if specalPrice, exists := c.pricingRules.SpecialPrices[sku]; exists {
			bundles := count / specalPrice.Count
			remainder := count % specalPrice.Count

			totalPrice += (bundles * specalPrice.SpecialTotal) + (remainder * c.pricingRules.Prices[sku])
		} else {
			totalPrice += count * c.pricingRules.Prices[sku]
		}
	}

	return totalPrice
}
