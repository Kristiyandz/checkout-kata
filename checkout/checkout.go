package checkout

import (
	"fmt"

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

func InitNewCheckout(pricingRules model.PricingRules) ICheckout {
	return &CheckoutBasket{
		pricingRules: pricingRules,
	}
}

func (c *CheckoutBasket) Scan(itemSKU string) {
	fmt.Printf("Scanning item: %s", itemSKU)
}

func (c *CheckoutBasket) GetTotalPrice() int {
	fmt.Println("Getting total value")
	return 0
}
