package core

import (
	"testing"
)

func TestObjLocatorInit(t *testing.T) {
	p := NewObjLocator(FieldPoint{2, 3}, NewEmptyPlace())
	if p == nil {
		t.Errorf("Expected not null pointer")
	}

	if p.location.x != 2 {
		t.Errorf("Coordinate x must be 2, but was %i", p.location.x)
	}

	if p.location.y != 3 {
		t.Errorf("Coordinate y must be 2, but was %i", p.location.y)
	}

	if p.obj.GetType() != Empty {
		t.Errorf("Obj type must be %i, but was %i", Empty, p.obj.GetType())
	}

}
