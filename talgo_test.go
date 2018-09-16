package talgo_test

import (
	"sort"
	"testing"

	"github.com/SebastienDorgan/talgo"
	"github.com/stretchr/testify/assert"
)

func GreaterThan(s sort.IntSlice, value int) talgo.Predicate {
	return func(i int) bool {
		return s[i] >= value
	}
}

type BankAccount struct {
	ID      string
	Balance int64
}

type BankAccounts []BankAccount

func (b BankAccounts) Len() int {
	return len(b)
}
func Test(t *testing.T) {
	s := sort.IntSlice{1, 2, 4}
	//Map Incr over s
	talgo.ForEach(s, func(i int) {
		s[i] += 3
	})
	assert.Equal(t, 4, s[0])
	assert.Equal(t, 5, s[1])
	assert.Equal(t, 7, s[2])

	//Find first element greater or equal to  5
	r := talgo.FindFirst(s, GreaterThan(s, 5))
	assert.Equal(t, 1, r)

	//Find last element greater or equal to 5
	r = talgo.FindLast(s, func(i int) bool {
		return s[i] >= 5
	})
	assert.Equal(t, 2, r)

	//Find last element greater or equal to 5
	l := talgo.FindAll(s, GreaterThan(s, 5))
	assert.Equal(t, 1, l[0])
	assert.Equal(t, 2, l[1])

	//Checks if at least one element is greater or equal to 5
	b := talgo.Any(s, GreaterThan(s, 5))
	assert.True(t, b)

	//Checks if at least one element is greater or equal to 0
	b = talgo.Any(s, talgo.Negate(GreaterThan(s, 0)))
	assert.False(t, b)

	//Checks if all elements are greater or equal to 5
	b = talgo.All(s, GreaterThan(s, 5))
	assert.False(t, b)

	//Reduce by adding all elements of the list to 3
	sum := 3
	talgo.ForEach(s, func(i int) {
		sum += s[i]
	})
	assert.Equal(t, 19, sum)

	//Multiply by 2 all elements of the list
	sf := sort.Float64Slice{2.5, 1.5, 3.5}
	talgo.ForEach(sf, func(i int) {
		sf[i] *= 2.
	})
	assert.Equal(t, 5., sf[0])
	assert.Equal(t, 3., sf[1])
	assert.Equal(t, 7., sf[2])

	//Find index of the min value
	min := func(i, j int) int {
		if sf[i] >= sf[j] {
			return j
		} else {
			return i
		}
	}

	minargs := talgo.Select(sf, min)

	assert.Equal(t, 1, minargs)

	//Find index of the max value
	max := func(i, j int) int {
		if sf[i] <= sf[j] {
			return j
		} else {
			return i
		}
	}

	maxargs := talgo.Select(sf, max)

	assert.Equal(t, 2, maxargs)

	//Separate odd and even elements
	ints := sort.IntSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
	even := []int{}
	odd := []int{}

	talgo.ForEach(ints, func(i int) {
		if ints[i]%2 == 0 {
			even = append(even, ints[i])
		} else {
			odd = append(odd, ints[i])
		}
	})

	//Find bank accounts in deficit
	bankAccounts := BankAccounts{
		BankAccount{ID: "0", Balance: 1500},
		BankAccount{ID: "1", Balance: -500},
		BankAccount{ID: "2", Balance: 0},
	}
	accountInDeficit := []string{}
	talgo.ForEach(bankAccounts, func(i int) {
		if bankAccounts[i].Balance < 0 {
			accountInDeficit = append(accountInDeficit, bankAccounts[i].ID)
		}
	})
	assert.Equal(t, "1", accountInDeficit[0])

}
