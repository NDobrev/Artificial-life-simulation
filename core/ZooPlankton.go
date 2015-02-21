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

func (zp *ZooPlankton) Do(f FieldBase, myLocation FieldPoint) {
	// try too eat "Grrrr fresh meat"
	// search for phyto plankton
	pp := f.GetAllWithTypeInSquare(GenFuncForObjType(LightSensitivePlanktonT), myLocation, 3)
	pp = append(pp, f.GetAllWithTypeInSquare(GenFuncForObjType(PhytoPlanktonT), myLocation, 3)...)
	if len(pp) > 0 {
		// eat with 80% rate
		if rand.Int()%100 < 50 {
			zp.energy += 2
			pos := rand.Int() % len(pp)
			f.RemoveFrom(pp[pos].location)
		}
	}
	if zp.TimeForDie() {
		f.RemoveFrom(myLocation)
		return
	}

	free := f.GetAllWithTypeInSquare(IsReplaceble, FieldPoint{myLocation.x, myLocation.y}, 3)
	if len(free) > 0 {
		if zp.TimeForReproduce() && zp.energy > 3 {
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
