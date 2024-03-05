package model

type Item struct {
	SKU       string
	UnitPrice int
}

type SpecialPrice struct {
	Count        int
	SpecialTotal int
}

type PricingRules struct {
	Prices        map[string]int
	SpecialPrices map[string]SpecialPrice
}
