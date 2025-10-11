package main

import "testing"

func TestHasExoplanet(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"665544554", false},
		{"FGFFCFFGG", true},
		{"MONOPLONOMONPLNOMPNOMP", false},
		{"FREECODECAMP", true},
		{"9AB98AB9BC98A", false},
		{"ZXXWYZXYWYXZEGZXWYZXYGEE", true},
	}

	for _, test := range tests {
		if result := hasExpoPlanet(test.input); result != test.want {
			t.Errorf("hasExpoPlanet(%s) = %t, want %t", test.input, result, test.want)
		}
	}
}
