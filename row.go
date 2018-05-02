package twenty48

type Row []Cell

// Reverse reverses all the data and returns a new row
func (cells Row) Reverse() Row {
	rowLength := len(cells)
	newCells := make(Row, rowLength)
	for i, j := 0, rowLength-1; i <= j; i, j = i+1, j-1 {
		newCells[i], newCells[j] = cells[j], cells[i]
	}
	return newCells
}

// Move moves all the non-zero cells to a particular direction
// and returns a new now
func (cells Row) Move(direction Direction) Row {
	length := len(cells)
	newCells := make(Row, length)

	rev := direction == DirectionRight || direction == DirectionDown
	if rev {
		cells = cells.Reverse()
	}

	// Move non-zero cells to the left
	pos := 0
	for _, c := range cells {
		if c.Value() > 0 {
			newCells[pos].SetValue(c.Value())
			pos++
		}
	}

	if rev {
		newCells = newCells.Reverse()
	}
	return newCells
}

// Merge moves and merges cell to the left and return the new list.
// It returns an error if no changes are made.
func (cells Row) Merge(direction Direction) Row {
	cells = cells.Move(direction)

	// we reverse in right and down case, because consecutive merges
	// are order sensitive
	// [2,2,2] ==Left=> [4,2,0]
	// [2,2,2] =Right=> [0,2,4]
	if direction == DirectionRight || direction == DirectionDown {
		cells = cells.Reverse()
	}

	// Merge, but ignore the last cell
	for k := 0; k < len(cells)-1; k++ {
		if cells[k].Value() == cells[k+1].Value() {
			cells[k].SetValue(cells[k].Value() + cells[k+1].Value())
			cells[k+1].SetValue(0)

			// skip the next one
			k++
		}
	}

	// reverse again to put in correct order
	if direction == DirectionRight || direction == DirectionDown {
		cells = cells.Reverse()
	}

	// remove any possible blank space after merging
	// TODO: Do we really need this?
	cells = cells.Move(direction)
	return cells
}
