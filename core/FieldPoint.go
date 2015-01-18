package core

type FieldPoint struct {
	x int
	y int
}

func (fp *FieldPoint) SetPoint(x int, y int) {
	fp.x = x
	fp.y = y
}

func NewFieldPoint(x int, y int) *FieldPoint {
	p := new(FieldPoint)
	p.x = x
	p.y = y
	return p
}
