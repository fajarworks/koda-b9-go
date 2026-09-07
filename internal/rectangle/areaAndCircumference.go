package rectangle

func AreaAndCircumOfRectangle(p int8, l int8) (area int16, circum int16) {
	area = int16(p) * int16(l)
	circum = 2 * (int16(p) + int16(l))
	return area, circum
}
