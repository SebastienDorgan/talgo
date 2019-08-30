package talgo

//Negate negates a predicate
func Negate(p Predicate) Predicate {
	return func(i int) bool {
		return !p(i)
	}
}

//Or create a new predicate that is the 'or' of predicate p1 and p2
func Or(p1 Predicate, p2 Predicate) Predicate {
	return func(i int) bool {
		return p1(i) || p2(i)
	}
}

//And create a new predicate that is the 'and' of predicate p1 and p2
func And(p1 Predicate, p2 Predicate) Predicate {
	return func(i int) bool {
		return p1(i) && p2(i)
	}
}

//Xor create a new predicate that is the 'xor' of predicate p1 and p2
func Xor(p1 Predicate, p2 Predicate) Predicate {
	return func(i int) bool {
		return p1(i) != p2(i)
	}
}
