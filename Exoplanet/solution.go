/*
Space Week Day 2: Exoplanet Search
Luminosity readings only comprise of characters 0-9 and A-Z that correspond to levels 0-35
*/
package main

import (
	"errors"
	"fmt"
)

const threshold = 0.8

func main() {
	fmt.Printf("Has Expo Planet: %t\n", hasExpoPlanet("665544554"))
	fmt.Printf("Has Expo Planet: %t\n", hasExpoPlanet("FGFFCFFGG"))
}

// hasExpoPlanet returns true if any luminosity reading is below the threshold
func hasExpoPlanet(readings string) bool {
	average := getAverageReading(readings)

	for _, reading := range readings {
		r, err := fromBase36(reading)

		if err != nil {
			continue
		}

		if float64(r) <= (average * threshold) {
			return true
		}
	}

	return false
}

func getAverageReading(readings string) float64 {
	sum := 0.0

	for _, reading := range readings {
		r, err := fromBase36(reading)

		if err != nil {
			continue
		}

		sum += float64(r)
	}

	return sum / float64(len(readings))
}

func fromBase36(r rune) (int, error) {
	if int(r) <= int('9') && int(r) >= int('0') {
		return int(r) - int('0'), nil
	}

	if int(r) >= int('A') && int(r) <= int('Z') {
		return int(r) - int('A') + 10, nil
	}

	return 0, errors.New("not a valid base 36 character")
}
