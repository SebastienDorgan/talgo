# talgo

**talgo** provides a set of algorithms (map, reduce, predicate, find...) that do not use reflection or empty interfaces

## Features
* Strong typing
* No overhead
* Map
* Reduce
* Predicate
* Find
* ...


## Dependencies
**talgo** uses Testify https://github.com/stretchr/testify

## Examples


```golang
package talgo_test

import (
	"sort"
	"testing"

	"github.com/SebastienDorgan/talgo"
	"github.com/stretchr/testify/assert"
)

func Incr(s sort.IntSlice, step int) talgo.Visitor {
	return func(i int) {
		s[i] = s[i] + step
	}
}

func GreaterThan(s sort.IntSlice, value int) talgo.Predicate {
	return func(i int) bool {
		return s[i] >= value
	}
}

func Acc(s sort.IntSlice, sum *int) talgo.Visitor {
	return func(i int) {
		*sum += s[i]
	}
}

func MultiplyBy(s sort.Float64Slice, x float64) talgo.Visitor {
	return func(i int) {
		s[i] = s[i] * x
	}
}

func Test(t *testing.T) {
	s := sort.IntSlice{1, 2, 4}
	//Map Incr over s
	talgo.Walk(s, Incr(s, 3))
	assert.Equal(t, 4, s[0])
	assert.Equal(t, 5, s[1])
	assert.Equal(t, 7, s[2])

	//Find first element greater or equal to  5
	r := talgo.FindFirst(s, GreaterThan(s, 5))
	assert.Equal(t, 1, r)

	//Find last element greater or equal to 5
	r = talgo.FindLast(s, GreaterThan(s, 5))
	assert.Equal(t, 2, r)

	//Checks if at least one element is greater or equal to 5
	b := talgo.Any(s, GreaterThan(s, 5))
	assert.True(t, b)

	//Checks if at least one element is greater or equal to 0
	b = talgo.Any(s, talgo.Negate(GreaterThan(s, 0)))
	assert.False(t, b)

	//Checks if all elements are greater or equal to 5
	b = talgo.All(s, GreaterThan(s, 5))
	assert.False(t, b)

	//Reduce
	sum := 3
	talgo.Walk(s, Acc(s, &sum))
	assert.Equal(t, 19, sum)

	sf := sort.Float64Slice{1.5, 2.5, 3.5}
	talgo.Walk(sf, MultiplyBy(sf, 2.))
	assert.Equal(t, 3., sf[0])
	assert.Equal(t, 5., sf[1])
	assert.Equal(t, 7., sf[2])
}

```
