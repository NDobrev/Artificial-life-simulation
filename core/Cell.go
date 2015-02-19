package core

import (
	"math/rand"
)

type Cell struct {
	age              int
	deadTime         int
	reproductionTime int
}

func NewCell(deadTime int, reproductionTime int) {
	result := new(Cell)
	result.deadTime = deadTime
	result.reproductionTime = reproductionTime
}

func (c *Cell) Do(f FieldBase, myLocation FieldPoint) {
	if c.TimeForReproduce() {
		free := f.GetAllWithTypeInSquare(IsReplaceble, FieldPoint{myLocation.x, myLocation.y}, 3)
		if len(free) != 0 {
			//expand on random
			pos := rand.Int() % len(free)
			npp := NewPhytoPlankton()
			f.AddObject(free[pos].location, npp)
		}
	}

	if c.TimeForDie() {
		f.RemoveFrom(myLocation)
	}
	c.age++
}

func (c *Cell) TimeForDie() bool {
	if c.age == c.deadTime {
		return true
	}
	return false
}

func (c *Cell) TimeForReproduce() bool {
	if c.age%c.reproductionTime == 0 {
		return true
	}
	return false
}

func (c *Cell) GetAge() int {
	return c.age
}
