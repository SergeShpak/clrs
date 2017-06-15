package sorting

type InsertionSorter struct{}

func (sorter InsertionSorter) Sort(
	sortables ISortables) ISortables {
	for i := 1; i < sortables.Len(); i++ {
		for k := i - 1; k >= 0 && sortables.Less(k+1, k); k-- {
			sortables.Swap(k+1, k)
		}
	}
	return sortables
}
