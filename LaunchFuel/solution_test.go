package main

import "testing"

func TestLaunchFuel(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{
			input: 50.0,
			want:  12.4,
		},
		{
			input: 500.0,
			want:  124.8,
		},
		{
			input: 243.0,
			want:  60.7,
		},
		{
			input: 11000.0,
			want:  2749.8,
		}, {
			input: 6214.0,
			want:  1553.4,
		},
	}

	for _, test := range tests {
		if result := launchFuel(test.input); result != test.want {
			t.Errorf("launchFuel(%.2f) = %f, want %.1f", test.input, result, test.want)
		}
	}
}
