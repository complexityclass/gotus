package doubly_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushFront(t *testing.T) {
	//given
	var dlist DList
	dlist.PushFront(6)
	dlist.PushFront(7)
	dlist.PushFront(8)
	var expected = []int{8, 7, 6}

	//when
	var result = listToIntSlice(dlist)

	//then
	assert.Equal(t, expected, result)
}

func TestPushBack(t *testing.T) {
	//given
	var dlist DList
	dlist.PushBack(6)
	dlist.PushBack(7)
	dlist.PushBack(8)
	var expected = []int{6, 7, 8}

	//when
	var result = listToIntSlice(dlist)

	//then
	assert.Equal(t, expected, result)
}

func TestFirstLast(t *testing.T) {
	//given
	var dlist DList
	dlist.PushFront(1)
	dlist.PushBack(2)
	dlist.PushBack(3)

	//when
	var first = dlist.First()
	var last = dlist.Last()

	//then
	assert.Equal(t, 1, first.Value().(int))
	assert.Equal(t, 3, last.Value().(int))
}

func TestRemove(t *testing.T) {
	//given
	var dlist DList
	var input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range input {
		dlist.PushBack(v)
	}

	var slice1 = []int{1, 2, 3, 5, 6, 7, 8, 9}
	var slice2 = []int{1, 2, 3, 5, 6, 8, 9}
	var slice3 = []int{2, 3, 5, 6, 8, 9}
	var slice4 = []int{2, 3, 5, 6, 8}

	//when
	var d1 = elemByValue(4, dlist)
	var d2 = elemByValue(7, dlist)
	var d3 = elemByValue(1, dlist)
	var d4 = elemByValue(9, dlist)

	d1.Remove()
	var res1 = listToIntSlice(dlist)
	d2.Remove()
	var res2 = listToIntSlice(dlist)
	d3.Remove()
	var res3 = listToIntSlice(dlist)
	d4.Remove()
	var res4 = listToIntSlice(dlist)

	//then
	assert.Equal(t, slice1, res1)
	assert.Equal(t, slice2, res2)
	assert.Equal(t, slice3, res3)
	assert.Equal(t, slice4, res4)
}

func listToIntSlice(dlist DList) []int {
	var result []int
	for node := dlist.First(); node != nil; node = node.Next() {
		var value = node.Value().(int)
		result = append(result, value)
	}

	return result
}

func elemByValue(value int, dlist DList) *DListItem {
	for node := dlist.First(); node != nil; node = node.Next() {
		if node.value.(int) == value {
			return node
		}
	}
	return nil
}
