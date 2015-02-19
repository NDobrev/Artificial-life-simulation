package core

import (
	"math/rand"
)

type PredatoryPlankton struct {
	Cell
	energy int
	speed  int
}

func NewPredatoryPlankton() *PredatoryPlankton {
	result := new(PredatoryPlankton)
	result.age = 0
	result.deadTime = 10
	result.reproductionTime = 5
	result.energy = 5
	result.speed = 7
	return result

}

func (pp *PredatoryPlankton) GetCopy() FieldObject {
	result := new(PredatoryPlankton)
	result.age = pp.age
	result.speed = pp.speed
	return result
}

func (pp *PredatoryPlankton) GetType() ObjType {
	return PredatoryPlanktonT
}

func (pp *PredatoryPlankton) Do(f FieldBase, myLocation FieldPoint) {
	// try too eat "Grrrr fresh meat"
	// search for phyto plankton
	targets := f.GetAllWithTypeInSquare(GenFuncForObjType(ZooPlanktonT), myLocation, pp.speed)
	if len(targets) > 0 {
		// eat with  rate
		if rand.Int()%100 < 88 {
			pos := rand.Int() % len(targets)
			pp.energy += targets[pos].obj.(*ZooPlankton).energy + 5
			f.RemoveFrom(targets[pos].location)
		}
	}
	if pp.TimeForDie() {
		f.RemoveFrom(myLocation)
		return
	}

	free := f.GetAllWithTypeInSquare(IsReplaceble, FieldPoint{myLocation.x, myLocation.y}, pp.speed)
	if len(free) > 0 {
		if pp.TimeForReproduce() && pp.energy > 5 {
			//expand on random
			pos := rand.Int() % len(free)
			npp := NewPredatoryPlankton()
			f.AddObject(free[pos].location, npp)
			pp.energy -= 5
		} else {
			pos := rand.Int() % len(free)
			f.MoveFromTo(myLocation, free[pos].location)
		}
	}

	pp.energy--
	pp.age++
}

func (pp *PredatoryPlankton) TimeForDie() bool {
	if pp.energy < 1 {
		return true
	}
	return pp.Cell.TimeForDie()
}
