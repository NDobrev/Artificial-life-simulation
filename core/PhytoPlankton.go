package core

import (
//"math/rand"
)

type PhytoPlankton struct {
	Cell
}

func NewPhytoPlankton() *PhytoPlankton {
	result := new(PhytoPlankton)
	result.age = 0
	result.deadTime = 10
	result.reproductionTime = 3
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

// noting spacial for now
