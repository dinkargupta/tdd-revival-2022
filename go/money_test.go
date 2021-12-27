package main

import "testing"

type Dollar struct {
	amount int
}

type Money struct {
	amount   int
	currency string
}

func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}

func TestMultiplication(t *testing.T) {
	fiver := Dollar{amount: 5}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got: [%d]", tenner.amount)
	}
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
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
