package twenty48

import "fmt"

type Grid []Row

func (rows Grid) Transpose() Grid {
	r := make(Grid, len(rows[0]))
	for x := range r {
		r[x] = make(Row, len(rows))
	}
	for y, cells := range rows {
		for x, e := range cells {
			r[x][y] = e
		}
	}
	return r
}

func (rows Grid) String() string {
	str := ""
	for _, cells := range rows {
		for _, e := range cells {
			str += fmt.Sprintf("%d", e.Value())
		}
	}
	return str
}
