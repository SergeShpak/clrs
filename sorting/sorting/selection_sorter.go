package sorting

type SelectionSorter struct{}

func (sorter SelectionSorter) Sort(sortables []ISortable, order Order) []ISortable {
	for i, _ := range sortables {
		swapIndex := i
		for j := i + 1; j < len(sortables); j++ {
			var comparison bool
			switch order {
			case Ascending:
				comparison = sortables[swapIndex].Compare(sortables[j]) < 0
			case Descending:
				comparison = sortables[swapIndex].Compare(sortables[j]) > 0
			default:
				panic("Sorting with the passed order has not been implemented yet")
			}
			if comparison {
				swapIndex = j
			}
		}
		if swapIndex != i {
			sortables[swapIndex], sortables[i] = sortables[i], sortables[swapIndex]
		}
	}
	return sortables
}
