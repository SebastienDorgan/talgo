package talgo

//Greater check if s[i] is greater or greater or equal to value
func Uint8Greater(s []uint8, value uint8, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func Uint8Lesser(s []uint8, value uint8, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func Uint8Equal(s []uint8, value uint8) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}
