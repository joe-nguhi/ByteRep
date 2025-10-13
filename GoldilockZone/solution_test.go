package main

import "testing"

func TestGoldilocksZone(t *testing.T) {
	tests := []struct {
		input float64
		want  [2]float64
	}{
		{
			input: 1,
			want: [2]float64{
				0.95,
				1.37,
			},
		},
		{
			input: 0.5,
			want: [2]float64{
				0.28,
				0.41,
			},
		},
		{
			input: 6,
			want: [2]float64{
				21.85,
				31.51,
			},
		},
		{
			input: 3.7,
			want: [2]float64{
				9.38,
				13.52,
			},
		},
		{
			input: 20,
			want: [2]float64{
				179.69,
				259.13,
			},
		},
	}

	for _, test := range tests {
		if result := goldilocksZone(test.input); result != test.want {
			t.Errorf("goldilocksZone(%f) = %v, want %v", test.input, result, test.want)
		}
	}
}
