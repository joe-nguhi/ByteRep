package main

import "testing"

func TestFindLandingSpot(t *testing.T) {
	tests := []struct {
		input [][]int
		want  [2]int
	}{
		{
			[][]int{
				{1, 0},
				{2, 0},
			},
			[2]int{0, 1},
		},
		{
			[][]int{
				{9, 0, 3},
				{7, 0, 4},
				{8, 0, 5},
			},
			[2]int{1, 1},
		},
		{
			[][]int{
				{1, 2, 1},
				{0, 0, 2},
				{3, 0, 0},
			},
			[2]int{2, 2},
		},
		{
			[][]int{
				{9, 6, 0, 8},
				{7, 1, 1, 0},
				{3, 0, 3, 9},
				{8, 6, 0, 9},
			},
			[2]int{2, 1},
		},
	}

	for _, test := range tests {
		if result := findLandingSpot(test.input); result != test.want {
			t.Errorf("findLandingSpot(%v) = %v, want %v", test.input, result, test.want)
		}
	}
}
