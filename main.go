package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Kristiyandz/checkout-kata/model"
)

func main() {
	var pricingRules model.PricingRules
	pricingRulesFile := "pricing_rules.json"

	file, err := os.Open(pricingRulesFile)
	if err != nil {
		log.Fatalf("Failed to open pricing rules file: %s", err)
	}
	defer file.Close()

	byteValye, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read pricing rules json: %s", err)
	}

	err = json.Unmarshal(byteValye, &pricingRules)

	if err != nil {
		log.Fatalf("Failed to unamrshal rules: %s", err)
	}

	fmt.Println(pricingRules)

	fmt.Println("Hello World!")
}
