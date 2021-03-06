package core

type LitPlace struct {
	lightIntensity int
}

func NewLitPlace(LightIntensity int) *LitPlace {
	result := new(LitPlace)
	result.lightIntensity = LightIntensity
	return result
}

func (lp *LitPlace) GetCopy() FieldObject {
	result := new(LitPlace)
	result.lightIntensity = lp.lightIntensity
	return result
}

func (lp *LitPlace) GetType() ObjType {
	return LitSpaceT
}
