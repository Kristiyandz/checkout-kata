package checkout_test

import (
	"testing"

	"github.com/Kristiyandz/checkout-kata/checkout"
	"github.com/Kristiyandz/checkout-kata/model"
)

func TestGetTotalPrices(t *testing.T) {
	testPricingRules := model.PricingRules{
		Prices: map[string]int{
			"A": 50,
			"B": 30,
			"C": 20,
			"D": 15,
		},
		SpecialPrices: map[string]model.SpecialPrice{
			"A": {
				Count:        3,
				SpecialTotal: 130,
			},
			"B": {
				Count:        2,
				SpecialTotal: 45,
			},
		},
	}

	newCheckout := checkout.InitNewCheckout(testPricingRules)

	newCheckout.Scan("A")
	newCheckout.Scan("B")
	newCheckout.Scan("A")
	newCheckout.Scan("A") // triggers special price because it is the second occurrence of A as per the rules
	newCheckout.Scan("B") // trigger special price because it is the second occurrence of B as per the rules
	newCheckout.Scan("C")

	expectedTotal := 130 + 45 + 20 // Special price for A and B + price for C
	actualTotal := 195             // this will become newCheckout.GetTotalPrice()

	if actualTotal != expectedTotal {
		t.Errorf("GetTotalPrice() = %d; want '%d'", actualTotal, expectedTotal)
	}
}
