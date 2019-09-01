package talgo

//Function defines a function that applies to the ith elements of a collection
type Function func(i int)

//Predicate defines a predicate that applies to the ith elements of a collection
type Predicate func(i int) bool

//Selector defines a selection function, a selector return i value or j value
type Selector func(i, j int) int

//Index use to reindex a collection
type Index func(i int) int

//ReverseIndex inverse the index of a collection
func ReverseIndex(len int) Index {
	return func(i int) int {
		return len - i - 1
	}
}

//ForEach apply m to each element of s
func ForEach(len int, f Function) {
	for i := 0; i < len; i++ {
		f(i)
	}
}

//ReverseForEach apply m to each element of s in reverse order
func ReverseForEach(len int, f Function) {
	for i := len - 1; i >= 0; i-- {
		f(i)
	}
}

func IndexedForEach(len int, index Index, f Function) {
	for i := 0; i < len; i++ {
		f(index(i))
	}
}

//Select select best element against selector s
func Select(len int, s Selector) int {
	selected := 0
	for i := 1; i < len; i++ {
		selected = s(selected, i)
	}
	return selected
}

//FindFirst find the first element of a collections that satisfies predicate p
func FindFirst(len int, p Predicate) int {
	for i := 0; i < len; i++ {
		if p(i) {
			return i
		}
	}
	return -1
}

//FindLast find the last element of a collections that satisfies predicate p
func FindLast(len int, p Predicate) int {
	for i := len - 1; i >= 0; i-- {
		if p(i) {
			return i
		}
	}
	return -1
}

//FindAll find all the elements of a collections that satisfies predicate p
func FindAll(len int, p Predicate) []int {
	var indexes []int
	for i := 0; i < len; i++ {
		if p(i) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

//CountItems counts number of items that satisfies predicate p
func CountItems(len int, p Predicate) int {
	cpt := 0
	for i := 0; i < len; i++ {
		if p(i) {
			cpt++
		}
	}
	return cpt
}

//Any checks if at least one element of the collection satisfies predicate p
func Any(len int, p Predicate) bool {
	i := FindFirst(len, p)
	return i >= 0
}

//None checks if at none element of the collection satisfies predicate p
func None(len int, p Predicate) bool {
	return !Any(len, p)
}

//All checks if all elements of the collection satisfy predicate p
func All(len int, p Predicate) bool {
	for i := 0; i < len; i++ {
		if !p(i) {
			return false
		}
	}
	return true
}
