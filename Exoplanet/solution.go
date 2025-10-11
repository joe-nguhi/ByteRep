package main

import (
	"errors"
)

const threshold = 0.8

func main() {
	//fmt.Printf("Has Expo Planet: %t\n", hasExpoPlanet("665544554"))
	//fmt.Printf("Has Expo Planet: %t\n", hasExpoPlanet("FGFFCFFGG"))
}

func hasExpoPlanet(readings string) bool {
	average := getAverageReading(readings)

	for _, reading := range readings {
		r, err := fromBase36(reading)

		if err != nil {
			continue
		}

		if float64(r) <= (float64(average) * threshold) {
			return true
		}
	}

	return false
}

func getAverageReading(readings string) int {
	sum := 0

	for _, reading := range readings {
		r, err := fromBase36(reading)

		if err != nil {
			continue
		}

		sum += r
	}

	return sum / len(readings)
}

func fromBase36(r rune) (int, error) {
	if int(r) < 58 && int(r) > 47 {
		return int(r) - 48, nil
	}

	if int(r) > 64 && int(r) < 91 {
		return int(r) - 55, nil
	}

	return 0, errors.New("not a valid base 36 character")
}
