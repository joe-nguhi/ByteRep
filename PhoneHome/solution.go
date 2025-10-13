/*
Space Week Day 3: Phone Home
Problem Statement: Given a communication route between your spaceship,
a series of satellites and your home planet, calculate the time it takes to send a message.
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	SPEED    float64 = 300_000.00
	SATDELAY float64 = 0.5
)

func main() {
	r := []int{54600000, 54600400}
	result := sendMessage(r)
	fmt.Printf("%v-\t%s\n", r, result)
}

func sendMessage(route []int) string {
	distance := 0

	for _, d := range route {
		distance += d
	}

	nsats := len(route) - 1

	result := float64(distance)/SPEED + (SATDELAY * float64(nsats))

	// Round to 4 decimal places properly
	rounded := math.Round(result*10000) / 10000

	s := fmt.Sprintf("%.4f", rounded)

	// Remove trailing zeros but keep at least one digit before decimal point

	s = strings.TrimRight(s, "0") // Remove trailing zeros
	s = strings.TrimRight(s, ".") // Remove trailing decimal point if present

	return s
}
