package core

import (
	"fmt"
)

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

	type ObjLoc struct {
		fo FieldObject
		po FieldPoint
	}

	first := GenRandomVector(lf.size)
	second := GenRandomVector(lf.size)
	lim := make([]ObjLoc, 0)
	for _, i := range first {
		for _, j := range second {
			if lf.matrix[i][j].GetType() > FirstDoable {
				lf.matrix[i][j].(DoableObject).Do(lf, FieldPoint{i, j})
				continue
			}
			if lf.matrix[i][j].GetType() == LimiterT {
				fmt.Println("Limiter!!")
				lim = append(lim, ObjLoc{lf.matrix[i][j], FieldPoint{i, j}})
				continue
			}
		}
	}

	for i := range lim {
		lim[i].fo.(*Limiter).Restruct(lf, lim[i].po)
	}
}

func (lf *LitField) ClearField() {
	for i := range lf.matrix {
		for j := range lf.matrix[i] {
			lf.matrix[i][j] = NewLitPlace(lf.FillFunc(i, j))
		}
	}
}
