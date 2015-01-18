package core

import (
	"math/rand"
)

type ZooPlankton struct {
	Cell
	energy int
}

func NewZooPlankton() *ZooPlankton {
	result := new(ZooPlankton)
	result.age = 0
	result.deadTime = 10
	result.reproductionTime = 3
	result.energy = 3
	return result

}

func (zp *ZooPlankton) GetCopy() FieldObject {
	result := new(ZooPlankton)
	result.age = zp.age
	return result
}

func (zp *ZooPlankton) GetType() ObjType {
	return ZooPlanktonT
}

func (zp *ZooPlankton) Do(f *Field, myLocation FieldPoint) {
	// try too eat "Grrrr fresh meat"
	// search for phyto plankton
	pp := f.GetAllWithTypeInSquare(PhytoPlanktonT, myLocation, 3)
	if len(pp) > 0 {
		// eat with 50% rate
		if rand.Int()%2 == 0 {
			zp.energy += 3
			pos := rand.Int() % len(pp)
			f.RemoveFrom(pp[pos].location)
		}
	}
	if zp.TimeForDie() {
		f.RemoveFrom(myLocation)
		return
	}

	free := f.GetAllWithTypeInSquare(Empty, FieldPoint{myLocation.x, myLocation.y}, 3)
	if len(free) > 0 {
		if zp.TimeForReproduce() && zp.energy > 5 {
			//expand on random
			pos := rand.Int() % len(free)
			npp := NewZooPlankton()
			f.AddObject(free[pos].location, npp)
			zp.energy--
		} else {
			pos := rand.Int() % len(free)
			f.MoveFromTo(myLocation, free[pos].location)
		}
	}

	zp.energy--
	zp.age++
}

func (zp *ZooPlankton) TimeForDie() bool {
	if zp.energy < 1 {
		return true
	}
	return zp.Cell.TimeForDie()
}
