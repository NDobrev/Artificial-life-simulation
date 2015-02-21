package core

import (
	"errors"
	"fmt"
)

type Field struct {
	matrix [][]FieldObject
	size   int
}

func NewField(size int) *Field {
	result := new(Field)
	result.size = size
	result.matrix = make([][]FieldObject, size)
	for i := range result.matrix {
		result.matrix[i] = make([]FieldObject, size)
	}

	filler := func() FieldObject {
		return NewEmptyPlace()
	}
	FillFieldWith(result.matrix, filler)
	return result
}

func (f *Field) AddObject(p FieldPoint, obj FieldObject) bool {
	if !f.checkCorectPoint(p) {
		return false
	}
	if !IsReplaceble(f.matrix[p.x][p.y]) {
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

	return f.GetAllWithTypeInSquare(GenFuncForObjType(t), FieldPoint{0, 0}, f.size)
}

func (f *Field) checkCorectPoint(p FieldPoint) bool {
	if p.x < 0 || p.y < 0 || p.x > f.size || p.y > f.size {
		return false
	}
	return true
}

func (f *Field) GetAllWithTypeInSquare(comp func(FieldObject) bool, center FieldPoint, size int) []*ObjLocator {
	result := make([]*ObjLocator, 0)

	topLeft := NewFieldPoint(center.x-size/2, center.y-size/2)
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
			if comp(f.matrix[i][j]) {
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
		f.matrix[from.x][from.y] = NewEmptyPlace()
		return true
	}

	return false
}

func (f *Field) RemoveFrom(from FieldPoint) {
	if !f.checkCorectPoint(from) {
		return
	}
	f.matrix[from.x][from.y] = NewEmptyPlace()
}

func (f *Field) ClearField() {
	for i := range f.matrix {
		for j := range f.matrix[i] {
			f.matrix[i][j] = NewEmptyPlace()
		}
	}
}

func (f *Field) OnTick() {
	for i := range f.matrix {
		for j := range f.matrix[i] {
			if f.matrix[i][j].GetType() > FirstDoable {
				f.matrix[i][j].(DoableObject).Do(f, FieldPoint{i, j})
			}
		}
	}
}

/*This suck but for now is ok */
func (f *Field) Print() {

	maping := func(ot ObjType) string {
		switch ot {
		case Empty:
			return "E"
		case LitSpaceT:
			return "L"

		case ZooPlanktonT:
			return "Z"
		case PhytoPlanktonT:
			return "P"
		default:
			return "Nan"
		}
	}
	fmt.Println("-------------")
	for i := range f.matrix {
		for j := range f.matrix[i] {
			fmt.Print(maping(f.matrix[i][j].GetType()))
		}
		fmt.Println("")
	}
	fmt.Println("-------------")
}
func clamp(x uint) uint {
	if x > 255 {
		x = 255
	}
	return x
}

func rgb(r, g, b uint) uint {
	return (clamp(r) << 16) | (clamp(g) << 8) | clamp(b)
}

func (f *Field) objRepresentation(obj FieldObject) uint {

	AgeMapingMutatable := func(age int, rep int) uint {
		result := rgb(0, uint(age*20), 255-uint(rep))
		return result
	}

	AgeMapingPhyto := func(age int) uint {
		return rgb(0, uint(age*10), 0)
	}
	AgeMapingZoo := func(age int) uint {
		return rgb(uint(age*10), 0, 0)
	}
	AgeMappingPred := func(age int) uint {
		return rgb(uint(age*10), 0, uint(age*10))
	}

	LightMaping := func(lightPower int) uint {
		return rgb(255, 255, uint(lightPower*5))
	}

	maping := func(ot ObjType) uint {
		switch ot {
		case Empty:
			return rgb(0, 0, 0)
		case LitSpaceT:
			return rgb(255, 160, 0)
		case RockT:
			return rgb(255, 255, 255)
		default:
			return 0
		}
	}

	switch obj.GetType() {
	case ZooPlanktonT:
		//fmt.Println("Z")
		return AgeMapingZoo(obj.(*ZooPlankton).GetAge())
	case PhytoPlanktonT:
		//fmt.Println("P")
		return AgeMapingPhyto(obj.(*PhytoPlankton).GetAge())
	case PredatoryPlanktonT:
		return AgeMappingPred(obj.(*PredatoryPlankton).GetAge())
	case LightSensitivePlanktonT:
		return AgeMapingMutatable(obj.(*LightSensitivePlankton).GetAge(), obj.(*LightSensitivePlankton).lightPreferenceMax+obj.(*LightSensitivePlankton).lightPreferenceMin)
	case LitSpaceT:
		return LightMaping(obj.(*LitPlace).lightIntensity)
	default:
		//fmt.Println("N")
		return maping(obj.GetType())

	}
	//fmt.Println("Lo6o")
	return 0
}

func (f *Field) ColorRepresentation(colors [][]uint) {

	for i := range f.matrix {
		for j := range f.matrix[i] {
			colors[i][j] = f.objRepresentation(f.matrix[i][j])
		}
	}
}
