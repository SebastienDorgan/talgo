package talgo

//Collection a type the satisfies Collection interface can be used by the routines of this package
type Collection interface {
	//Len is the number of elements of the collection
	Len() int
}

//Function defines a function that applies to the ith elements of a collection
type Function func(i int)

//Predicate defines a predicate that applies to the ith elements of a collection
type Predicate func(i int) bool

//Selector defines a selection function, a selector return i value or j value
type Selector func(i, j int) int

//Index use to reindex a collection
type Index func(i int) int

//ReverseIndex inverse the index of a collection
func ReverseIndex(c Collection) Index {
	l := c.Len()
	return func(i int) int {
		return l - i - 1
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

//FindFirst find the first element of a collections that satisfies predicate p
func FindFirst(c Collection, p Predicate) int {
	for i := 0; i < c.Len(); i++ {
		if p(i) {
			return i
		}
	}
	return -1
}

//FindLast find the last element of a collections that satisfies predicate p
func FindLast(c Collection, p Predicate) int {
	for i := c.Len() - 1; i >= 0; i-- {
		if p(i) {
			return i
		}
	}
	return -1
}

//FindAll find all the elements of a collections that satisfies predicate p
func FindAll(c Collection, p Predicate) IntSlice {
	var indexes IntSlice
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

//Any checks if at least one element of the collection satisfies predicate p
func Any(c Collection, p Predicate) bool {
	i := FindFirst(c, p)
	println(i)
	return i >= 0
}

//None checks if at none element of the collection satisfies predicate p
func None(c Collection, p Predicate) bool {
	return !Any(c, p)
}

//All checks if all elements of the collection satisfy predicate p
func All(c Collection, p Predicate) bool {
	for i := 0; i < c.Len(); i++ {
		if !p(i) {
			return false
		}
	}
	return true
}
