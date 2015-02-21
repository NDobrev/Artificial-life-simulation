package core

import (
	//	"fmt"
	"math/rand"
)

type PhytoPlankton struct {
	Cell
}

func NewPhytoPlankton() *PhytoPlankton {
	result := new(PhytoPlankton)
	result.age = 0
	result.deadTime = 10
	result.reproductionTime = 5
	return result
}

func (pp *PhytoPlankton) GetCopy() FieldObject {
	result := new(PhytoPlankton)
	result.age = pp.age
	return result
}

func (pp *PhytoPlankton) GetType() ObjType {
	return PhytoPlanktonT
}

/*Recproduction habit are base on light on field*/
type LightSensitivePlankton struct {
	Cell
	lightPreferenceMin int
	lightPreferenceMax int
	mutationRate       int // one of how many cell will be mutate
}

func NewLightSensitivePlankton() *LightSensitivePlankton {
	result := new(LightSensitivePlankton)
	result.age = 0
	result.deadTime = 30
	result.reproductionTime = 9
	result.lightPreferenceMin = 0
	result.lightPreferenceMax = 5
	result.mutationRate = 15
	return result
}

func (lsp *LightSensitivePlankton) GenChild() *LightSensitivePlankton {
	result := new(LightSensitivePlankton)
	result.age = 0
	result.deadTime = lsp.deadTime
	result.reproductionTime = lsp.reproductionTime
	result.lightPreferenceMin = lsp.lightPreferenceMin
	result.lightPreferenceMax = lsp.lightPreferenceMax
	result.mutationRate = lsp.mutationRate
	return result
}

func (lsp *LightSensitivePlankton) Do(f FieldBase, myLocation FieldPoint) {

	comp := func(fo FieldObject) bool {
		if fo.GetType() == LitSpaceT && InInterval(lsp.lightPreferenceMin, lsp.lightPreferenceMax, fo.(*LitPlace).lightIntensity) {
			return true
		}
		return false
	}
	free := f.GetAllWithTypeInSquare(comp, myLocation, 3)

	if len(free) != 0 {
		//expand on random
		pos := rand.Int() % len(free)
		npp := lsp.GenChild()
		f.AddObject(free[pos].location, npp)
	}

	if lsp.TimeForDie() || lsp.lightPreferenceMax-lsp.lightPreferenceMin > 7 {
		f.RemoveFrom(myLocation)
	}
	lsp.Mutate()
	lsp.age++

}

func (lsp *LightSensitivePlankton) GetCopy() FieldObject {
	result := new(LightSensitivePlankton)
	// this is not valid
	return result
}

func (lsp *LightSensitivePlankton) GetType() ObjType {
	return LightSensitivePlanktonT
}

func (lsp *LightSensitivePlankton) Mutate() {

	if 0 == rand.Int()%lsp.mutationRate {
		lsp.lightPreferenceMax += rand.Int() % 2
		lsp.lightPreferenceMax -= rand.Int() % 2
		lsp.lightPreferenceMin += rand.Int() % 2
		lsp.lightPreferenceMin -= rand.Int() % 2
		lsp.deadTime += rand.Int() % 2
		lsp.deadTime -= rand.Int() % 2
		lsp.reproductionTime += rand.Int() % 2
		lsp.reproductionTime -= rand.Int() % 2
	}
}

// noting spacial for now
