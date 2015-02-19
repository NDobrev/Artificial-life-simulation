package core

import (
	"sync"
)

type Rock struct {
	sync.Mutex
}

func NewRock() *Rock {
	return new(Rock)
}

func (r *Rock) GetCopy() FieldObject {
	return NewRock()
}

func (r *Rock) GetType() ObjType {
	return RockT
}
