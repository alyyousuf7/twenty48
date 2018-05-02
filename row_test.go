package twenty48_test

import (
	"testing"

	"github.com/alyyousuf7/twenty48"
)

func TestRowReverse(t *testing.T) {
	// prepare data
	row := twenty48.Row{1, 2, 3, 4, 5}
	rowLength := len(row)

	// do reverse
	newRow := row.Reverse()

	// check if the length is correct
	if rowLength != len(newRow) {
		t.Error("Reversed row length does not match row length")
	}

	// check if the values are as expected
	for k, r := range row {
		if r.Value() != newRow[rowLength-k-1].Value() {
			t.Errorf("row[%d]: %d == %d value did not match", k, r.Value(), newRow[rowLength-k-1].Value())
			break
		}
	}
}

type MoveTestCase struct {
	row         twenty48.Row
	direction   twenty48.Direction
	expectedrow twenty48.Row
}

func testRowMoveCase(t *testing.T, testcase MoveTestCase) {
	row := testcase.row
	direction := testcase.direction

	// prepare data
	rowLength := len(row)

	// move
	newRow := row.Move(direction)

	// check if the length is correct
	if rowLength != len(newRow) {
		t.Error("Moved row length does not match row length")
	}

	// check if the values are as expected
	for k, r := range newRow {
		if r.Value() != testcase.expectedrow[k].Value() {
			t.Errorf("expected row %v but got %v", testcase.expectedrow, newRow)
			break
		}
	}
}

func TestRowMove(t *testing.T) {
	cases := []MoveTestCase{
		{ // extreme move
			row:         twenty48.Row{0, 0, 0, 0, 2},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{2, 0, 0, 0, 0},
		},
		{
			row:         twenty48.Row{2, 0, 0, 0, 0},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{0, 0, 0, 0, 2},
		},
		{ // no move with zeroes
			row:         twenty48.Row{4, 2, 0, 0, 0},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{4, 2, 0, 0, 0},
		},
		{
			row:         twenty48.Row{0, 0, 0, 4, 2},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{0, 0, 0, 4, 2},
		},
		{ // no moves, because no zeroes
			row:         twenty48.Row{1024, 512, 256, 128, 64},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{1024, 512, 256, 128, 64},
		},
		{
			row:         twenty48.Row{1024, 512, 256, 128, 64},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{1024, 512, 256, 128, 64},
		},
		{ // move multiple
			row:         twenty48.Row{0, 4, 0, 0, 2},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{4, 2, 0, 0, 0},
		},
		{
			row:         twenty48.Row{0, 4, 0, 0, 2},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{0, 0, 0, 4, 2},
		},
		{ // move a few
			row:         twenty48.Row{4, 0, 0, 0, 2},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{4, 2, 0, 0, 0},
		},
		{
			row:         twenty48.Row{4, 0, 0, 0, 2},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{0, 0, 0, 4, 2},
		},
	}

	// Multiple similar testcases: up == left, down == right
	for _, c := range cases {
		testRowMoveCase(t, c)

		if c.direction == twenty48.DirectionLeft {
			c.direction = twenty48.DirectionUp
			testRowMoveCase(t, c)
		}

		if c.direction == twenty48.DirectionRight {
			c.direction = twenty48.DirectionDown
			testRowMoveCase(t, c)
		}
	}
}

type MergeTestCase struct {
	row         twenty48.Row
	direction   twenty48.Direction
	expectedrow twenty48.Row
}

func testRowMergeCase(t *testing.T, testcase MergeTestCase) {
	row := testcase.row
	direction := testcase.direction

	// prepare data
	rowLength := len(row)

	// merge
	newRow := row.Merge(direction)

	// check if the length is correct
	if rowLength != len(newRow) {
		t.Error("Merged row length does not match row length")
	}

	// check if the values are as expected
	for k, r := range newRow {
		if r.Value() != testcase.expectedrow[k].Value() {
			t.Errorf("expected row %v but got %v", testcase.expectedrow, newRow)
			break
		}
	}
}

func TestRowMerge(t *testing.T) {
	cases := []MergeTestCase{
		{ // no merges, only moves
			row:         twenty48.Row{0, 0, 0, 0, 2},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{2, 0, 0, 0, 0},
		},
		{
			row:         twenty48.Row{2, 0, 0, 0, 0},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{0, 0, 0, 0, 2},
		},
		{ // adjacent merges, no moves
			row:         twenty48.Row{2, 2, 0, 0, 0},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{4, 0, 0, 0, 0},
		},
		{
			row:         twenty48.Row{0, 0, 0, 2, 2},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{0, 0, 0, 0, 4},
		},
		{ // adjacent merges, with complete move
			row:         twenty48.Row{0, 0, 0, 2, 2},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{4, 0, 0, 0, 0},
		},
		{
			row:         twenty48.Row{2, 2, 0, 0, 0},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{0, 0, 0, 0, 4},
		},
		{ // first two should merge, not the other
			row:         twenty48.Row{2, 2, 2, 0, 0},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{4, 2, 0, 0, 0},
		},
		{
			row:         twenty48.Row{0, 0, 2, 2, 2},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{0, 0, 0, 2, 4},
		},
		{ // multiple merges
			row:         twenty48.Row{2, 2, 2, 2, 0},
			direction:   twenty48.DirectionLeft,
			expectedrow: twenty48.Row{4, 4, 0, 0, 0},
		},
		{
			row:         twenty48.Row{0, 2, 2, 2, 2},
			direction:   twenty48.DirectionRight,
			expectedrow: twenty48.Row{0, 0, 0, 4, 4},
		},
	}

	// Multiple similar testcases: up == left, down == right
	for _, c := range cases {
		testRowMergeCase(t, c)

		if c.direction == twenty48.DirectionLeft {
			c.direction = twenty48.DirectionUp
			testRowMergeCase(t, c)
		}

		if c.direction == twenty48.DirectionRight {
			c.direction = twenty48.DirectionDown
			testRowMergeCase(t, c)
		}
	}
}
