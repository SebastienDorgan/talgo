package talgo

//Greater check if s[i] is greater or greater or equal to value
func Uint16Greater(s []uint16, value uint16, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func Uint16Lesser(s []uint16, value uint16, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func Uint16Equal(s []uint16, value uint16) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}
