package twenty48_test

import (
	"testing"

	"github.com/alyyousuf7/twenty48"
)

func TestGridTranspose(t *testing.T) {
	// prepare data
	grid := twenty48.Grid{
		twenty48.Row{1, 2, 3, 4, 5},
		twenty48.Row{6, 7, 8, 9, 10},
		twenty48.Row{11, 12, 13, 14, 15},
		twenty48.Row{16, 17, 18, 19, 20},
	}

	expectedGrid := twenty48.Grid{
		twenty48.Row{1, 6, 11, 16},
		twenty48.Row{2, 7, 12, 17},
		twenty48.Row{3, 8, 13, 18},
		twenty48.Row{4, 9, 14, 19},
		twenty48.Row{5, 10, 15, 20},
	}

	// do reverse
	newGrid := grid.Transpose()

	// check if the values are as expected
	for kr, r := range expectedGrid {
		for kc, c := range r {
			if c.Value() != newGrid[kr][kc].Value() {
				t.Errorf("Transposed grid did not return expected result")
				break
			}
		}
	}
}

func TestGridString(t *testing.T) {
	// prepare data
	grid := twenty48.Grid{
		twenty48.Row{1, 2, 3, 4, 5, 6},
		twenty48.Row{7, 8, 9, 10, 11, 12},
	}

	expectedStr := "123456789101112"

	str := grid.String()

	if expectedStr != str {
		t.Errorf("Expected %s, but got %s", expectedStr, str)
	}
}
