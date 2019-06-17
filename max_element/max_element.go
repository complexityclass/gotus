package max_element

import "reflect"

// FindMax is a function to find max element in slice
func FindMax(sl interface{}, less func(i, j interface{}) bool) (interface{}, bool) {
	var slType = reflect.ValueOf(sl)
	var size = slType.Len()

	if size == 0 {
		return nil, false
	}

	slice := make([]interface{}, size)
	for i := 0; i < size; i++ {
		slice[i] = slType.Index(i).Interface()
	}

	var currentMax = slice[0]
	for _, elem := range slice {
		if less(currentMax, elem) {
			currentMax = elem
		}
	}

	return currentMax, true
}
