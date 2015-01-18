package core

type EmptyPlace struct {
}

func NewEmptyPlace() *EmptyPlace {
	return new(EmptyPlace)
}

func (ep *EmptyPlace) GetCopy() FieldObject {
	return new(EmptyPlace)
}

func (ep *EmptyPlace) GetType() ObjType {
	return Empty
}
