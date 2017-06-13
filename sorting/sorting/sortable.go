package sorting

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

func ConvertToISortable(sortables []IntSortable) []ISortable {
	result := make([]ISortable, len(sortables))
	for i, el := range sortables {
		result[i] = ISortable(el)
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
