package main

/* SpaceWeek Day 1:Stellar Classification */

import "fmt"

func main() {

	fmt.Printf("Temp:\t5778  Star Class: %s\n", classification(5778))
	fmt.Printf("Temp:\t2400  Star Class: %s\n", classification(2400))
	fmt.Printf("Temp:\t9999  Star Class: %s\n", classification(9999))
}

// classification Given the surface temperature of a star in Kelvin(K),return its stellar classication
func classification(temperature int) string {
	switch {
	case temperature >= 30000:
		return ("O")
	case temperature >= 10000:
		return "B"
	case temperature >= 7500:
		return "A"
	case temperature >= 6000:
		return "F"
	case temperature >= 5200:
		return "G"
	case temperature >= 3700:
		return "K"
	case temperature >= 0:
		return "M"
	default:
		return "Unclassified"
	}
}
