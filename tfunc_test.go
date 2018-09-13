package tfunc_test

import (
	"testing"

	"github.com/SebastienDorgan/tfunc"
	"github.com/stretchr/testify/assert"
)

type IntList []int

func (s IntList) Len() int {
	return len(s)
}

func Incr(s IntList, step int) tfunc.Visitor {
	return func(i int) {
		s[i] = s[i] + step
	}
}

func GreaterThan(s IntList, value int) tfunc.Predicate {
	return func(i int) bool {
		return s[i] >= value
	}
}

func Acc(s IntList, sum *int) tfunc.Visitor {
	return func(i int) {
		*sum += s[i]
	}
}

func Test(t *testing.T) {
	s := IntList{1, 2, 4}
	//Map Incr over s
	tfunc.Walk(s, Incr(s, 3))
	assert.Equal(t, 4, s[0])
	assert.Equal(t, 5, s[1])
	assert.Equal(t, 7, s[2])

	//Find first element greater or equal to  5
	r := tfunc.FindFirst(s, GreaterThan(s, 5))
	assert.Equal(t, 1, r)

	//Find last element greater or equal to 5
	r = tfunc.FindLast(s, GreaterThan(s, 5))
	assert.Equal(t, 2, r)

	//Checks if at least one element is greater or equal to 5
	b := tfunc.Any(s, GreaterThan(s, 5))
	assert.True(t, b)

	//Checks if at least one element is greater or equal to 0
	b = tfunc.Any(s, tfunc.Negate(GreaterThan(s, 0)))
	assert.False(t, b)

	//Checks if all elements are greater or equal to 5
	b = tfunc.All(s, GreaterThan(s, 5))
	assert.False(t, b)

	//Reduce
	sum := 3
	tfunc.Walk(s, Acc(s, &sum))
	assert.Equal(t, 19, sum)
}
