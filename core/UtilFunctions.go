package core

func InInterval(leftBound int, rightBound int, value int) bool {
	if leftBound < value && value < rightBound {
		return true
	}
	return false
}
