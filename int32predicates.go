package talgo

//Greater check if s[i] is greater or greater or equal to value
func Int32Greater(s []int32, value int32, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func Int32Lesser(s []int32, value int32, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func Int32Equal(s []int32, value int32) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}
