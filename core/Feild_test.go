package core

import (
	"testing"
)

func TestInitField(t *testing.T) {
	f := NewField(5)

	if f == nil {
		t.Errorf("Expected not null pointer")
	}

	if len(f.matrix) != 5 {
		t.Errorf("Expected field with size 5 but was %i", len(f.matrix))
	}
}

func TestAddObjInField(t *testing.T) {
	f := NewField(5)

	f.AddObject(FieldPoint{2, 2}, NewPhytoPlankton())

	if f.matrix[2][2].GetType() != PhytoPlanktonT {
		t.Errorf("Type of field must be %i but we found %i", f.matrix[2][2].GetType(), PhytoPlanktonT)
	}
}

func TestLookAtWithCorrectPoint(t *testing.T) {
	f := NewField(5)

	f.AddObject(FieldPoint{2, 2}, NewPhytoPlankton())

	result, err := f.LookAt(FieldPoint{2, 2})

	if err != nil {
		t.Errorf("We dont expected error", err)
	}

	if result.GetType() != PhytoPlanktonT {
		t.Errorf("Type of field must be %i but we found %i", f.matrix[2][2].GetType(), PhytoPlanktonT)
	}

}

func TestLookAtWithWrongPoint(t *testing.T) {
	f := NewField(5)
	p := FieldPoint{10, 2}
	f.AddObject(p, NewPhytoPlankton())

	_, err := f.LookAt(p)

	if err == nil {
		t.Errorf("We expected error")
	}
}

func TestMoveFromToWithCorrectPoint(t *testing.T) {
	f := NewField(5)
	p := FieldPoint{2, 2}
	f.AddObject(p, NewPhytoPlankton())
	np := FieldPoint{1, 1}
	f.MoveFromTo(p, np)

	if f.matrix[p.x][p.y].GetType() != Empty {
		t.Errorf("Type of field must be %i but we found %i", f.matrix[p.x][p.y].GetType(), PhytoPlanktonT)
	}

	if f.matrix[np.x][np.y].GetType() != PhytoPlanktonT {
		t.Errorf("Type of field must be %i but we found %i", f.matrix[np.x][np.y].GetType(), PhytoPlanktonT)
	}

}

func TestMoveFromToWithWrongPoint(t *testing.T) {
	f := NewField(5)
	p := FieldPoint{2, 2}
	f.AddObject(p, NewPhytoPlankton())
	np := FieldPoint{1, 10}
	f.MoveFromTo(p, np)

	if f.matrix[p.x][p.y].GetType() != PhytoPlanktonT {
		t.Errorf("Type of field must be %i but we found %i", f.matrix[p.x][p.y].GetType(), PhytoPlanktonT)
	}

}

func TestGetAllWithTypeInSquare(t *testing.T) {
	/*
		we have
		0 0 0
		0 3 0
		0 0 3
		and we want to check for point (1,1) with size 3  result must be 2
	*/
	f := NewField(3)
	f.AddObject(FieldPoint{2, 2}, NewPhytoPlankton())
	f.AddObject(FieldPoint{1, 1}, NewPhytoPlankton())

	result := f.GetAllWithTypeInSquare(PhytoPlanktonT, FieldPoint{2, 2}, 3)
	if len(result) != 2 {
		t.Errorf("We expected 2 cells with type PhytoPlanktonT, but we found", len(result))
	}

	result = f.GetAllWithTypeInSquare(PhytoPlanktonT, FieldPoint{0, 0}, 3)

	if len(result) != 1 {
		t.Errorf("We expected 1 cells with type PhytoPlanktonT, but we found", len(result))
	}
}
