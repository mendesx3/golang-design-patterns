package main

import "fmt"

type TransportationMode interface {
	TravelTime(int) int
}

type carMode struct{}

func (c *carMode) TravelTime(distance int) int {
	return distance / 60
}

type BikeMode struct{}

func (b *BikeMode) TravelTime(distance int) int {
	return distance / 10
}

type PublicTransportationMode struct{}

func (p *PublicTransportationMode) TravelTime(distance int) int {
	return distance / 5
}

func main() {
	distance := 100

	var mode TransportationMode
	mode = &carMode{}
	travelTime := mode.TravelTime(distance)
	fmt.Printf("Car travel time: %d\n", travelTime)

	mode = &BikeMode{}
	travelTime = mode.TravelTime(distance)
	fmt.Printf("Bike travel time: %d\n", travelTime)

	mode = &PublicTransportationMode{}
	travelTime = mode.TravelTime(distance)
	fmt.Printf("Public transportation travel time: %d\n", travelTime)

}
