package talgo

//Greater check if s[i] is greater or greater or equal to value
func UintGreater(s []uint, value uint, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func UintLesser(s []uint, value uint, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func UintEqual(s []uint, value uint) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}
