package core

import (
	"testing"
)

func TestFieldPointInit(t *testing.T) {
	p := NewFieldPoint(1, 2)
	if p == nil {
		t.Errorf("Expected not null pointer")
	}

	if p.x != 1 {
		t.Errorf("Coordinate x must be 1, but was %i", p.location.x)
	}

	if p.y != 2 {
		t.Errorf("Coordinate y must be 2, but was %i", p.location.y)
	}
}

func TestFieldPointSet(t *testing.T) {
	p := NewFieldPoint(1, 2)
	p.SetPoint(4, 5)

	if p.x != 4 {
		t.Errorf("Coordinate x must be 4, but was %i", p.location.x)
	}

	if p.y != 5 {
		t.Errorf("Coordinate y must be 5, but was %i", p.location.y)
	}
}
