package talgo

//Collection a type the satisfies Collection interface can be used by the routines of this package
type Collection interface {
	//Len is the numbre of elements of the collection
	Len() int
}

//Function defines a function that applies to the ith elements of a collection
type Function func(i int)

//Predicate defines a predicate that applies to the ith elements of a collection
type Predicate func(i int) bool

//Selector defines a selection function, a selector return i value or j value
type Selector func(i, j int) int

//Negate negates a predicate
func Negate(p Predicate) Predicate {
	return func(i int) bool {
		return !p(i)
	}
}

//Or create a new predicte that is the 'or' of predicate p1 and p2
func Or(p1 Predicate, p2 Predicate) Predicate {
	return func(i int) bool {
		return p1(i) || p2(i)
	}
}

//And create a new predicte that is the 'and' of predicate p1 and p2
func And(p1 Predicate, p2 Predicate) Predicate {
	return func(i int) bool {
		return p1(i) && p2(i)
	}
}

//Xor create a new predicte that is the 'xor' of predicate p1 and p2
func Xor(p1 Predicate, p2 Predicate) Predicate {
	return func(i int) bool {
		return p1(i) != p2(i)
	}
}

//ForEach apply m to each element of s
func ForEach(c Collection, f Function) {
	for i := 0; i < c.Len(); i++ {
		f(i)
	}
}

//ReverseForEach apply m to each element of s in reverse order
func ReverseForEach(c Collection, f Function) {
	for i := c.Len() - 1; i >= 0; i-- {
		f(i)
	}
}

//Select select best element against selector s
func Select(c Collection, s Selector) int {
	selected := 0
	for i := 1; i < c.Len(); i++ {
		selected = s(selected, i)
	}
	return selected
}

//FindFirst find the first element of a series that satisfies predicate p
func FindFirst(c Collection, p Predicate) int {
	for i := 0; i < c.Len(); i++ {
		if p(i) {
			return i
		}
	}
	return -1
}

//FindLast find the last element of a series that satisfies predicate p
func FindLast(c Collection, p Predicate) int {
	for i := c.Len() - 1; i >= 0; i-- {
		if p(i) {
			return i
		}
	}
	return -1
}

//FindAll find all the elements of a series that satisfies predicate p
func FindAll(c Collection, p Predicate) []int {
	indexes := []int{}
	for i := 0; i < c.Len(); i++ {
		if p(i) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

//CountItems counts number of items that satisfies predicate p
func CountItems(c Collection, p Predicate) int {
	cpt := 0
	for i := 0; i < c.Len(); i++ {
		if p(i) {
			cpt++
		}
	}
	return cpt
}

//Any checks if at least one element of the serie satisfies predicate p
func Any(c Collection, p Predicate) bool {
	return FindFirst(c, p) >= 0
}

//None checks if at none element of the serie satisfies predicate p
func None(c Collection, p Predicate) bool {
	return !Any(c, p)
}

//All checks if all elements of the serie satisfy predicate p
func All(c Collection, p Predicate) bool {
	for i := 0; i < c.Len(); i++ {
		if !p(i) {
			return false
		}
	}
	return true
}
