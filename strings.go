package talgo

import "strings"

//Greater check if s[i] is greater or greater or equal to value
func GreaterString(s []string, value string, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return strings.Compare(s[i], value) >= 0
		}
		return strings.Compare(s[i], value) > 0
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func LesserString(s []string, value string, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return strings.Compare(s[i], value) <= 0
		}
		return strings.Compare(s[i], value) < 0
	}
}

//Equal check if s[i] is equal to value
func EqualString(s []string, value string) Predicate {
	return func(i int) bool {
		return strings.Compare(s[i], value) == 0
	}
}
