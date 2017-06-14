package sorting

import (
	"reflect"
)

type ISortable interface {
	Compare(ISortable) int
}

type IntSortable struct {
	Value int
}

func (this IntSortable) Compare(that ISortable) int {
	thatIntSortable, ok := that.(IntSortable)
	if !ok {
		panic("An attempt to compare two Sortables of different types")
	}
	return this.Value - thatIntSortable.Value
}

func ToISortable(objects interface{}) []ISortable {
	if reflect.TypeOf(objects).Kind() != reflect.Slice {
		panic("Object passed to ToISortable function must be a slice")
	}
	sortables := reflect.ValueOf(objects)
	result := make([]ISortable, sortables.Len())
	var ok bool
	for i := 0; i < sortables.Len(); i++ {
		result[i], ok = sortables.Index(i).Interface().(ISortable)
		if !ok {
			panic("Attempt to cast objects that do not implement ISortable to ISortable")
		}
	}
	return result
}

func ISortableToIntSortable(genericSortables []ISortable) []IntSortable {
	result := make([]IntSortable, len(genericSortables))
	for i, el := range genericSortables {
		conversionResult, ok := el.(IntSortable)
		if !ok {
			panic("An attempt to cast an ISortable of another type to IntSortable")
		}
		result[i] = conversionResult
	}
	return result
}
