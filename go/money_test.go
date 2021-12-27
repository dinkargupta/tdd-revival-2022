package main

import "testing"

type Money struct {
	amount   int
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func TestMultiplicationInDollar(t *testing.T) {
	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := fiveDollars.Times(2)
	if tenDollars.amount != 10 {
		t.Errorf("Expected 10, got: [%d]", tenDollars.amount)
	}
	if tenDollars.currency != "USD" {
		t.Errorf("Expected 10, got: [%s]", tenDollars.currency)
	}
}

func TestMultiplicationInEuro(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	twentyEuro := tenEuros.Times(2)

	if twentyEuro.amount != 20 {
		t.Errorf("Expected 10, got: [%d]", twentyEuro.amount)
	}

	if twentyEuro.currency != "EUR" {
		t.Errorf("Expected EUR, got: [%s]", twentyEuro.currency)
	}
}
