package core

func IsPlankton(fo FieldObject) bool {
	if fo.GetType() > FirstPlankton && fo.GetType() < LastPlankton {
		return true
	}
	return false
}
