package main

import (
	"fmt"
	"strings"
)

const (
	SPEED    float64 = 300_000.00
	SATDELAY float64 = 0.5
)

func main() {
	r := []int{384400, 384400}
	fmt.Printf("%v-\t%s\n", r, sendMessage(r))
}

func sendMessage(route []int) string {

	distance := 0

	for _, d := range route {
		distance += d
	}

	nsats := len(route) - 1

	// Calculate the result with full precision first
	result := float64(distance)/SPEED + (SATDELAY * float64(nsats))

	// Print the raw result and rounded result for debugging
	fmt.Printf("Raw result: %.10f, Rounded: %.4f\n", result, result)

	// Format with exactly 4 decimal places but remove trailing zeros
	s := fmt.Sprintf("%.4f", result)

	// Remove trailing zeros but keep at least one digit after decimal point if needed
	if strings.Contains(s, ".") {
		s = strings.TrimRight(s, "0")
		s = strings.TrimRight(s, ".")
	}

	return s
}
