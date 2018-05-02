package twenty48

import (
	"fmt"
	"math"
	"math/rand"
)

var (
	ErrNoEmptyCell    = fmt.Errorf("No empty cell")
	ErrNoMovePossible = fmt.Errorf("No move possible")
)

type Board struct {
	width  int
	height int
	cell   Grid
	moved  bool // moved in last move
}

func NewBoard(width, height int) *Board {
	grid := make(Grid, height)

	for i := 0; i < height; i++ {
		grid[i] = make(Row, width)
	}

	return &Board{
		width:  width,
		height: height,
		cell:   grid,
		moved:  true,
	}
}

func (b *Board) Size() (int, int) {
	return b.width, b.height
}

func (b *Board) At(x, y int) *Cell {
	return &b.cell[y][x]
}

func (b *Board) emptyCell() (*Cell, error) {
	cells := []*Cell{}
	for x := 0; x < b.width; x++ {
		for y := 0; y < b.height; y++ {
			if b.At(x, y).Empty() {
				cells = append(cells, b.At(x, y))
			}
		}
	}

	if len(cells) == 0 {
		return nil, ErrNoEmptyCell
	}

	return cells[rand.Int()%len(cells)], nil
}

func (b *Board) NewMove() {
	if b.moved {
		b.addNewCell()
	}
	b.moved = false
}

func (b *Board) addNewCell() error {
	cell, err := b.emptyCell()
	if err != nil {
		return err
	}

	randVal := int(math.Pow(2.0, float64(rand.Int()%3)+1))
	cell.SetValue(randVal)

	return nil
}

func (b *Board) moveVertical(direction Direction) error {
	if direction != DirectionUp && direction != DirectionDown {
		return nil
	}

	b.moved = false
	state := b.cell.String()

	b.cell = b.cell.Transpose()
	for x := 0; x < b.width; x++ {
		row := Row(b.cell[x])
		b.cell[x] = row.Merge(direction)
	}
	b.cell = b.cell.Transpose()

	if state != b.cell.String() {
		b.moved = true
	}

	return nil
}

func (b *Board) moveHorizontal(direction Direction) error {
	if direction != DirectionUp && direction != DirectionDown {
		return nil
	}

	b.moved = false
	state := b.cell.String()

	for y := 0; y < b.height; y++ {
		row := Row(b.cell[y])
		b.cell[y] = row.Merge(direction)
	}

	if state != b.cell.String() {
		b.moved = true
	}

	return nil
}

func (b *Board) MoveUp() error {
	return b.moveVertical(DirectionUp)
}

func (b *Board) MoveDown() error {
	return b.moveVertical(DirectionDown)
}

func (b *Board) MoveLeft() error {
	return b.moveHorizontal(DirectionUp)
}

func (b *Board) MoveRight() error {
	return b.moveHorizontal(DirectionDown)
}
