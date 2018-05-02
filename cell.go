package twenty48

type Direction int

const (
	DirectionLeft Direction = iota
	DirectionRight
	DirectionUp
	DirectionDown
)

type Cell int

func (c *Cell) Empty() bool {
	return *c == 0
}

func (c *Cell) SetValue(v int) {
	*c = Cell(v)
}

func (c *Cell) Value() int {
	return int(*c)
}
