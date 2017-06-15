package sorting

type SelectionSorter struct{}

func (sorter SelectionSorter) Sort(sortables ISortables) ISortables {
	for i := 0; i < sortables.Len(); i++ {
		swapIndex := i
		for j := i + 1; j < sortables.Len(); j++ {
			if sortables.Less(j, swapIndex) {
				swapIndex = j
			}
		}
		if swapIndex != i {
			sortables.Swap(swapIndex, i)
		}
	}
	return sortables
}
