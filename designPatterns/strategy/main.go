package main

import "fmt"

type OperationStrategy interface {
	Execute(int, int) int
}

type AdditionStrategy struct{}

func (AdditionStrategy) Execute(a, b int) int {
	return a + b
}

type SubtractionStrategy struct{}

func (s *SubtractionStrategy) Execute(a, b int) int {
	return a - b
}

type Calculator struct {
	strategy OperationStrategy
}

func (c *Calculator) SetStrategy(strategy OperationStrategy) {
	c.strategy = strategy
}

func (c *Calculator) Calculate(a, b int) int {
	if c.strategy == nil {
		panic("A estratégia não foi definida")
	}
	return c.strategy.Execute(a, b)
}

func main() {
	calculator := &Calculator{}

	// Define a estratégia de adição
	calculator.SetStrategy(&AdditionStrategy{})
	result := calculator.Calculate(5, 3)
	fmt.Println("5 + 3 =", result)

	// Define a estratégia de subtração
	calculator.SetStrategy(&SubtractionStrategy{})
	result = calculator.Calculate(5, 3)
	fmt.Println("5 - 3 =", result)
}
