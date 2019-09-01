package talgo

//Greater check if s[i] is greater or greater or equal to value
func Int64Greater(s []int64, value int64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func Int64Lesser(s []int64, value int64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func Int64Equal(s []int64, value int64) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}
