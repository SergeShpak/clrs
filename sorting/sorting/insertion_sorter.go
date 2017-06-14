package sorting

type InsertionSorter struct {
}

func (sorter InsertionSorter) Sort(sortables []ISortable, order Order) []ISortable {
	for i := 1; i < len(sortables); i++ {
		el := sortables[i]
		for j := 0; j <= i-1; j++ {
			var comparison bool
			switch order {
			case Ascending:
				comparison = el.Compare(sortables[j]) < 0
			case Descending:
				comparison = el.Compare(sortables[j]) > 0
			default:
				panic("Sorting with the passed order has not been implemented yet")
			}
			if comparison {
				for k := i - 1; k >= j; k-- {
					sortables[k+1] = sortables[k]
				}
				sortables[j] = el
				break
			}
		}
	}
	return sortables
}
