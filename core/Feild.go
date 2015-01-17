package core

type FieldPoint struct {
	x int
	y int
}

type ObjLocator struct {
	lacation FieldPoint
	obj      *FieldObject
}

type Field struct {
	matrix [][]*FieldObject
	size   int
}

func NewField(size int) *Field {
	result := new(Field)
	result.size = size
	result.matrix = make([][]*FieldObject, size)

	// init matrix field
	for i := range result.matrix {
		result.matrix[i] = make([]*FieldObject, size)
		for j := range result.matrix[i] {
			var e EmptyPlace
			j = e.GetCopy()
		}
	}

	return result
}

func (f *Field) AddObject(x int, y int, obj *FieldObject) bool {
	if f.matrix[x][y].GetType() == Empty {
		return false
	}
	f.matrix[x][y] = obj
}

func (f *Field) LookAt(x int, y int) *FieldObject {
	return f.matrix[x][y]
}

func (f *Field) GetAllWithType(t ObjType) []*ObjLocator {
	result := make([]*ObjLocator)
	for i := range f.matrix {
		for j := range i {
			if j.GetType() == t {
				result = append(result, j)
			}
		}
	}
	return result
}

func (f *Field) MoveFromTo(from FieldPoint, to FieldPoint) bool {
	if f.matrix[to.x][to.y] == Empty {
		f.matrix[to.x][to.y] = f.matrix[from.x][from.y]
		var e EmptyPlace
		f.matrix[from.x][from.y] = e.GetCopy()
	}
}

func (f *Field) RemoveFrom(from FieldPoint) {
	var e EmptyPlace
	f.matrix[from.x][from.y] = e.GetCopy()
}

func (f *Field) ClearField(){
	var e EmptyPlace
	for i := range f.matrix {
		for j := i {
			i = e.GetCopy()
		}
	}
}




