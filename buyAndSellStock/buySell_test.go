package main

import (
	"testing"
)

func Testbuyandsell(t *testing.T) {
	test := []struct {
		name    string
		arr     []int
		expectX int
		expectY int
	}{
		{
			name:    "Case with multiple increasing order",
			arr:     []int{7, 6, 2, 4, 3, 9},
			expectX: 7,
			expectY: 9,
		},
		{
			name:    "Case with decreasing order",
			arr:     []int{7, 6, 2, 4, 3, 1},
			expectX: 0,
			expectY: 0,
		},
	}

	for _, t1 := range test {
		t.Run(t1.name, func(t *testing.T) {
			x, y := buyandsell(t1.arr)
			if x != t1.expectX || y != t1.expectY {
				t.Errorf("buyandsell(%v) = (%d, %d); want (%d, %d)", t1.arr, x, y, t1.expectX, t1.expectY)

			}
		})

	}

}
