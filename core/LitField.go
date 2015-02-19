package core

import ()

type LitField struct {
	lightIntensity int
	Field
	FillFunc func(int, int) int
}

func NewLitField(size int, LightIntensity int) *LitField {
	result := new(LitField)
	result.size = size
	result.matrix = make([][]FieldObject, size)
	result.lightIntensity = LightIntensity
	for i := range result.matrix {
		result.matrix[i] = make([]FieldObject, size)
	}

	filler := func() FieldObject {
		return NewLitPlace(LightIntensity)
	}
	FillFieldWith(result.matrix, filler)
	return result
}

func NewLitFieldWithFillWithFunction(size int, FillFunc func(int, int) int) *LitField {
	result := new(LitField)
	result.size = size
	result.matrix = make([][]FieldObject, size)
	result.lightIntensity = -1
	result.FillFunc = FillFunc
	for i := range result.matrix {
		result.matrix[i] = make([]FieldObject, size)
	}

	for i := range result.matrix {
		for j := range result.matrix[i] {
			result.matrix[i][j] = NewLitPlace(FillFunc(i, j))
		}
	}
	return result
}

func (lf *LitField) RemoveFrom(from FieldPoint) {
	if !lf.checkCorectPoint(from) {
		return
	}
	if lf.lightIntensity > 0 {
		lf.matrix[from.x][from.y] = NewLitPlace(lf.lightIntensity)
	} else {
		lf.matrix[from.x][from.y] = NewLitPlace(lf.FillFunc(from.x, from.y))
	}
}

func (lf *LitField) OnTick() {
	for i := range lf.matrix {
		for j := range lf.matrix[i] {
			if lf.matrix[i][j].GetType() > FirstDoable {
				lf.matrix[i][j].(DoableObject).Do(lf, FieldPoint{i, j})
			}
		}
	}
}

func (lf *LitField) ClearField() {
	for i := range lf.matrix {
		for j := range lf.matrix[i] {
			lf.matrix[i][j] = NewLitPlace(lf.FillFunc(i, j))
		}
	}
}
