package talgo

//Greater check if s[i] is greater or greater or equal to value
func IntGreater(s []int, value int, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func IntLesser(s []int, value int, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func IntEqual(s []int, value int) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}
