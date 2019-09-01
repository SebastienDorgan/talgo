package talgo

//Greater check if s[i] is greater or greater or equal to value
func Int16Greater(s []int16, value int16, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func Int16Lesser(s []int16, value int16, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func Int16Equal(s []int16, value int16) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}
