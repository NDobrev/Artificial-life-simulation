package core

import (
	"errors"
	"fmt"
)

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

type ObjLocator struct {
	location FieldPoint
	obj      FieldObject
}

func NewObjLocator(p FieldPoint, obj FieldObject) *ObjLocator {
	o := new(ObjLocator)
	o.location = FieldPoint{p.x, p.y}
	o.obj = obj
	return o
}

type Field struct {
	matrix [][]FieldObject
	size   int
}

func NewField(size int) *Field {
	result := new(Field)
	result.size = size
	result.matrix = make([][]FieldObject, size)

	// init matrix field
	for i := range result.matrix {
		result.matrix[i] = make([]FieldObject, size)
		for j := range result.matrix[i] {
			var e EmptyPlace
			result.matrix[i][j] = e.GetCopy()
		}
	}
	return result
}

func (f *Field) AddObject(p FieldPoint, obj FieldObject) bool {
	if !f.checkCorectPoint(p) {
		return false
	}
	if f.matrix[p.x][p.y].GetType() != Empty {
		return false
	}
	f.matrix[p.x][p.y] = obj
	return true
}

func (f *Field) LookAt(p FieldPoint) (FieldObject, error) {
	if !f.checkCorectPoint(p) {
		return nil, errors.New("invalid location")
	}
	return f.matrix[p.x][p.y], nil
}

func (f *Field) GetAllWithType(t ObjType) []*ObjLocator {
	return f.GetAllWithTypeInSquare(t, FieldPoint{0, 0}, f.size)
}

func (f *Field) checkCorectPoint(p FieldPoint) bool {
	if p.x < 0 || p.y < 0 || p.x > f.size || p.y > f.size {
		return false
	}
	return true
}

func (f *Field) GetAllWithTypeInSquare(t ObjType, topLeft FieldPoint, size int) []*ObjLocator {
	result := make([]*ObjLocator, 0)

	min := func(x int, y int) int {
		if x > y {
			return y
		}
		return x
	}
	max := func(x int, y int) int {
		if x > y {
			return x
		}
		return y
	}
	fromX := max(0, topLeft.x)
	fromY := max(0, topLeft.y)
	toX := min(f.size, topLeft.x+size)
	toY := min(f.size, topLeft.y+size)
	for i := fromX; i < toX; i++ {
		for j := fromY; j < toY; j++ {
			if f.matrix[i][j].GetType() == t {
				o := NewObjLocator(FieldPoint{i, j}, f.matrix[i][j])
				result = append(result, o)
			}
		}
	}
	return result
}

func (f *Field) MoveFromTo(from FieldPoint, to FieldPoint) bool {
	if !f.checkCorectPoint(from) || !f.checkCorectPoint(to) {
		return false
	}

	if f.matrix[to.x][to.y].GetType() == Empty {
		f.matrix[to.x][to.y] = f.matrix[from.x][from.y]
		var e EmptyPlace
		f.matrix[from.x][from.y] = e.GetCopy()
		return true
	}

	return false
}

func (f *Field) RemoveFrom(from FieldPoint) {
	var e EmptyPlace
	if !f.checkCorectPoint(from) {
		return
	}
	f.matrix[from.x][from.y] = e.GetCopy()
}

func (f *Field) ClearField() {
	var e EmptyPlace
	for i := range f.matrix {
		for j := range f.matrix[i] {
			f.matrix[i][j] = e.GetCopy()
		}
	}
}

func (f *Field) OnTick() {
	for i := range f.matrix {
		for j := range f.matrix[i] {
			if f.matrix[i][j].GetType() > Doable {
				f.matrix[i][j].(DoableObject).Do(f, FieldPoint{i, j})
			}
		}
	}
}

func (f *Field) Print() {
	for i := range f.matrix {
		for j := range f.matrix[i] {
			fmt.Print(f.matrix[i][j].GetType())
		}
		fmt.Println("")
	}
}
