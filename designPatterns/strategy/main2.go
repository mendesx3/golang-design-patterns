package main

import "fmt"

type Item struct {
	Name  string
	Price float64
}

type ShoppingCart struct {
	items    []Item
	strategy PriceCalculatorStrategy
}

type PriceCalculatorStrategy interface {
	CalculatePrice([]Item) float64
}

type RegularPriceCalculation struct{}

func (r *RegularPriceCalculation) CalculatePrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price
	}
	return total
}

type DiscountPriceCalculation struct{}

func (d *DiscountPriceCalculation) CalculatePrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price
	}
	return total * 0.9 // Aplica um desconto de 10%
}

func main() {
	items := []Item{
		{Name: "Laptop", Price: 1000.0},
		{Name: "Mouse", Price: 20.0},
		{Name: "Keyboard", Price: 50.0},
	}

	// Crie um carrinho de compras com a estratégia de preço regular
	cart := ShoppingCart{
		items:    items,
		strategy: &RegularPriceCalculation{},
	}

	// Calcule e imprima o preço total
	totalRegular := cart.strategy.CalculatePrice(cart.items)
	fmt.Printf("Preço total (sem desconto): $%.2f\n", totalRegular)

	// Mude para a estratégia de preço com desconto
	cart.strategy = &DiscountPriceCalculation{}

	// Calcule e imprima o preço total com desconto
	totalDiscount := cart.strategy.CalculatePrice(cart.items)
	fmt.Printf("Preço total (com desconto): $%.2f\n", totalDiscount)
}
