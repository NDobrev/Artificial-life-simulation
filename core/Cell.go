package core

type PhytoPlankton struct {
	age int
}

func NewPhytoPlankton() *PhytoPlankton {
	result := new(PhytoPlankton)
	result.age = 0
	return result
}

func (pp *PhytoPlankton) GetCopy() *FieldObject {
	result := new(PhytoPlankton)
	result.age = pp.age
	return result
}

func (pp *PhytoPlankton) GetType() *ObjType {
	return Plankton
}

func (pp *PhytoPlankton) Do(f *Field) {
	if pp.TimeForReproduce() {

	}

	if pp.TimeForDie() {

	}
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
