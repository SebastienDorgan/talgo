package talgo

//Greater check if s[i] is greater or greater or equal to value
func Float64Greater(s []float64, value float64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func Float64Lesser(s []float64, value float64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func Float64Equal(s []float64, value float64) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}
