package tfunc_test

import (
	"testing"

	"github.com/CS-SI/tfunc"
)

type IntSeries []int

func (s IntSeries) Len() int {
	return len(s)
}

func Incr(s IntSeries, step int) tfunc.Visitor {
	return func(i int) {
		s[i] = s[i] + step
	}
}

func GreaterThan(s IntSeries, value int) tfunc.Predicate {
	return func(i int) bool {
		return s[i] >= value
	}
}

func Accumulator(s IntSeries, sum *int) tfunc.Visitor {
	return func(i int) {
		*sum += s[i]
	}
}

func Test(t *testing.T) {
	s := IntSeries{1, 2, 4}
	//Map
	tfunc.Walk(s, Incr(s, 3))
	for _, v := range s {
		println(v)
	}

	r := tfunc.FindFirst(s, GreaterThan(s, 5))
	println(r)

	r = tfunc.FindLast(s, GreaterThan(s, 5))
	println(r)

	b := tfunc.Any(s, GreaterThan(s, 5))
	println(b)

	b = tfunc.Any(s, tfunc.Negate(GreaterThan(s, 0)))
	println(b)

	b = tfunc.All(s, GreaterThan(s, 5))
	println(b)

	//Reduce
	sum := 0
	tfunc.Walk(s, Accumulator(s, &sum))
	println(sum)
}
