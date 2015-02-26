package core

import (
	//"fmt"
	"math/rand"
)

func IsPlankton(fo FieldObject) bool {
	if fo.GetType() > FirstPlankton && fo.GetType() < LastPlankton {
		return true
	}
	return false
}

func IsReplaceble(fo FieldObject) bool {
	if fo.GetType() > FirstReplaceble && fo.GetType() < LastReplaceble {
		return true
	}
	return false
}

func GenFuncForObjType(ot ObjType) func(FieldObject) bool {
	return func(fo FieldObject) bool {
		if fo.GetType() == ot {
			return true
		}
		return false
	}
}

func AllType(fo FieldObject) bool {
	return true
}

func GenRandomVector(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i
	}
	for i := 0; i < size; i++ {
		r := rand.Int() % size
		temp := result[i]
		result[i] = result[r]
		result[r] = temp
	}
	return result
}
