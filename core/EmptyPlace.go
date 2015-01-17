package core

type EmptyPlace struct {
}

func (EmptyPlace *ep) GetCopy() *EmptyPlace {
	return new(EmptyPlace)
}

func (EmptyPlace *ep) GetType() ObjType {
	return Empty
}
