package checkout_test

import (
	"fmt"
	"testing"

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

	fmt.Println(testPricingRules)

	// @TODO init new checkout
	// @TODO scan items

	expectedTotal := 0
	actualTotal := 0

	if actualTotal != expectedTotal {
		t.Errorf("GetTotalPrice() = %d; want '%d'", actualTotal, expectedTotal)
	}
}
