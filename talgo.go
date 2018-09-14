package talgo

//Collection a type the satisfies Collection interface can be used by the routines of this package
type Collection interface {
	//Len is the numbre of elements of the collection
	Len() int
}

//Visitor defines a function that applies to the ith elements of a collection
type Visitor func(i int)

//Predicate defines a predicate that applies to the ith elements of a collection
type Predicate func(i int) bool

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

//FindFirst find the first element of a series that satisfies predicate p
func FindFirst(s Collection, p Predicate) int {
	for i := 0; i < s.Len(); i++ {
		if p(i) {
			return i
		}
	}
	return -1
}

//FindLast find the last element of a series that satisfies predicate p
func FindLast(s Collection, p Predicate) int {
	for i := s.Len() - 1; i >= 0; i-- {
		if p(i) {
			return i
		}
	}
	return -1
}

//FindAll find all the elements of a series that satisfies predicate p
func FindAll(s Collection, p Predicate) []int {
	indexes := []int{}
	for i := s.Len() - 1; i >= 0; i-- {
		if p(i) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

//Any checks if at least one element of the serie match predicate p
func Any(s Collection, p Predicate) bool {
	return FindFirst(s, p) >= 0
}

//All checks if all elements of the serie match predicate p
func All(s Collection, p Predicate) bool {
	return !Any(s, p)
}
