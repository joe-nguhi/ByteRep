package main

import (
	"fmt"
	"log"
	"time"
)

const (
	DateRef    = "2000-01-06"
	DateFormat = "2006-01-02"
	FullCycle  = 28
)

func main() {
	d := "2012-10-21"
	fmt.Printf("Moon Phase on %s: %s\n", d, moonPhase(d))
}

func moonPhase(day string) string {

	d, err := time.Parse(DateFormat, day)
	if err != nil {
		log.Fatalf("Error parsing date: %v", err)
	}

	ref, err := time.Parse(DateFormat, DateRef)

	if err != nil {
		log.Fatalf("Error parsing date: %v", err)
	}

	daysDiff := d.Sub(ref).Hours() / 24.0

	dayOfCycle := int(daysDiff) % FullCycle

	switch {
	case dayOfCycle < 7:
		return "New"
	case dayOfCycle < 14:
		return "Waxing"
	case dayOfCycle < 21:
		return "Full"
	case dayOfCycle < 28:
		return "Waning"
	default:
		return "Unknown"
	}
}
