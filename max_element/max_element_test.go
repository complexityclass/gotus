package max_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptySlice(t *testing.T) {
	//given
	var sl []string

	//when
	var elem, ok = FindMax(sl, func(i, j interface{}) bool {
		return len(i.(string)) < len(j.(string))
	})

	//then
	assert.Nil(t, elem)
	assert.False(t, ok)
}

func TestMaxInteger(t *testing.T) {
	//given
	numbers := []int{1, 5, 19, 3, 46, 2}

	//when
	var elem, ok = FindMax(numbers, func(i, j interface{}) bool {
		return i.(int) < j.(int)
	})

	//then
	assert.True(t, ok)
	assert.Equal(t, 46, elem)
}

func TestMaxCustomType(t *testing.T) {
	//given
	type Host struct {
		name string
		rps  int
	}

	hosts := []Host{
		{"first", 40},
		{"second", 60},
		{"third", 20},
	}

	//when
	var elem, ok = FindMax(hosts, func(i, j interface{}) bool {
		return i.(Host).rps < j.(Host).rps
	})

	assert.True(t, ok)
	assert.Equal(t, "second", elem.(Host).name)
}
