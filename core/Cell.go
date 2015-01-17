package core

import (
	"fmt"
	"math/rand"
)

type PhytoPlankton struct {
	age int
}

func NewPhytoPlankton() *PhytoPlankton {
	result := new(PhytoPlankton)
	result.age = 0
	return result
}

func (pp *PhytoPlankton) GetCopy() FieldObject {
	result := new(PhytoPlankton)
	result.age = pp.age
	return result
}

func (pp *PhytoPlankton) GetType() ObjType {
	return Plankton
}

func (pp *PhytoPlankton) Do(f *Field, myLocation FieldPoint) {
	if pp.TimeForReproduce() {
		free := f.GetAllWithTypeInSquare(Empty, FieldPoint{myLocation.x - 1, myLocation.y - 1}, 3)
		if len(free) != 0 {
			//expand on random
			pos := rand.Int() % len(free)
			npp := NewPhytoPlankton()
			f.AddObject(free[pos].location, npp)
		}
	}

	if pp.TimeForDie() {
		fmt.Println(myLocation.x, myLocation.y, "die")
		f.RemoveFrom(myLocation)
	}
	pp.age++
}

func (pp *PhytoPlankton) TimeForDie() bool {
	if pp.age == 10 {
		return true
	}
	return false
}

func (pp *PhytoPlankton) TimeForReproduce() bool {
	if pp.age%3 == 0 {
		return true
	}
	return false
}
