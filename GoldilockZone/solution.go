/*
Space Week Day 5: Goldilocks Zone
Problem statement: Given the mass of a star, find the start and end of the Goldilocks Zone in Astronomical Units.
*/
package main

import (
	"fmt"
	"math"
)

const (
	LF = 3.5  // Luminosity Factor
	GS = 0.95 // Goldilocks Zone Start
	GE = 1.37 //  Goldilocks Zone End
)

func main() {
	fmt.Printf("%.2f\n", goldilocksZone(1))
}

func goldilocksZone(mass float64) [2]float64 {
	starLum := math.Pow(mass, LF)

	goldiStarts := GS * math.Sqrt(starLum)
	goldiEnds := GE * math.Sqrt(starLum)

	return [2]float64{
		roundTo(goldiStarts, 2),
		roundTo(goldiEnds, 2),
	}
}

func roundTo(n float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(n*factor) / factor
}
