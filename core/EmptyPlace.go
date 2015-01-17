package core

type EmptyPlace struct {
}

func (ep *EmptyPlace) GetCopy() FieldObject {
	return new(EmptyPlace)
}

func (ep *EmptyPlace) GetType() ObjType {
	return Empty
}
