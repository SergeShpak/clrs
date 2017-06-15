package sorting

import (
	"reflect"
)

type ISortables interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

/*
func ToISortables(objects interface{}) ISortables {
	if reflect.TypeOf(objects).Kind() != reflect.Slice {
		panic("Object passed to ToISortable function must be a slice")
	}
	sortables := reflect.ValueOf(objects)
	result := make([]IntSortable, sortables.Len())
	var ok bool
	for i := 0; i < sortables.Len(); i++ {
		result[i], ok = sortables.Index(i).Interface().(ISortable)
		if !ok {
			panic("Attempt to cast objects that do not implement ISortable to ISortable")
		}
	}
	return result
}
*/

func ToISortables(object interface{}) ISortables {
	if reflect.TypeOf(object).Kind() != reflect.Slice {
		panic("Object passed to ToISortable function must be a slice")
	}
	sortables := reflect.ValueOf(object).Interface()
	return sortables.(ISortables)
}

/*
func ToIntSortable(genericSortables ISortables) IntSortables {
	result := make(IntSortables, len(genericSortables))
	for i, el := range genericSortables {
		conversionResult, ok := el.(IntSortable)
		if !ok {
			panic("An attempt to cast an ISortable of another type to IntSortable")
		}
		result[i] = conversionResult
	}
	return result
}
*/

func ToIntSortables(genericSortables ISortables) IntSortables {
	return genericSortables.(IntSortables)
}

type IntSortables []int

func (this IntSortables) Len() int {
	return len(this)
}

func (this IntSortables) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this IntSortables) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func CreateIntSortables(ints ...int) IntSortables {
	result := make([]int, len(ints))
	for i, val := range ints {
		result[i] = val
	}
	return result
}

func (slice IntSortables) Copy() IntSortables {
	result := make([]int, slice.Len())
	for i, val := range slice {
		result[i] = val
	}
	return result
}
