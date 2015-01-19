package core

import (
//"fmt"
)

type LitField struct {
	lightIntensity int
	Field
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

func (lf *LitField) RemoveFrom(from FieldPoint) {
	if !lf.checkCorectPoint(from) {
		return
	}
	lf.matrix[from.x][from.y] = NewLitPlace(lf.lightIntensity)
}

func (lf *LitField) OnTick() {
	for i := range lf.matrix {
		for j := range lf.matrix[i] {
			if lf.matrix[i][j].GetType() > Doable {
				lf.matrix[i][j].(DoableObject).Do(lf, FieldPoint{i, j})
			}
		}
	}
}
