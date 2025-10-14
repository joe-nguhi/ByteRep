/*
Space Week Day 7: Launch Fuel
Problem Statement: Given the mass of a payload in kilograms, calculate the amount of fuel require to launch to orbit.
Note: Fuel itself has mass
*/

package main

import (
	"fmt"
	"math"
)

const WeightFuelRatio float64 = 1.0 / 5.0

func main() {
	fmt.Printf("Fuel Required: %.2f\n", launchFuel(243))
}

func launchFuel(payload float64) float64 {

	var calculate func(float64) float64

	calculate = func(payload float64) float64 {
		if payload < (1.0 / WeightFuelRatio) {
			return payload * WeightFuelRatio
		}

		fuel := payload * WeightFuelRatio
		newPayload := payload + fuel
		diff := newPayload - payload

		result := fuel + calculate(diff)

		return result
	}

	result := calculate(payload)

	return roundTo(result, 1)
}

func roundTo(n float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(n*factor) / factor
}
