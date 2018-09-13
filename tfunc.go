package tfunc

//Visitor defines a visitor function
type Visitor func(int)

//Collection defines a series
type Collection interface {
	Len() int
}

//Predicate defines a predicate function
type Predicate func(int) bool

//Negate negates a predictae
func Negate(p Predicate) Predicate {
	return func(i int) bool {
		return !p(i)
	}
}

//Walk apply m to each element of s
func Walk(s Collection, m Visitor) {
	for i := 0; i < s.Len(); i++ {
		m(i)
	}
}

//FindFirst find the first element of a series that matches predicate p
func FindFirst(s Collection, p Predicate) int {
	for i := 0; i < s.Len(); i++ {
		if p(i) {
			return i
		}
	}
	return -1
}

//FindLast find the lats element of a series that matches predicate p
func FindLast(s Collection, p Predicate) int {
	for i := s.Len() - 1; i >= 0; i-- {
		if p(i) {
			return i
		}
	}
	return -1
}

//Any checks if at least one element of the serie match predicate p
func Any(s Collection, p Predicate) bool {
	return FindFirst(s, p) >= 0
}

//All checks if all elements of the serie match predicate p
func All(s Collection, p Predicate) bool {
	return !Any(s, p)
}
