package sorting

type InsertionSorter struct {
}

func (sorter InsertionSorter) Sort(sortables []ISortable) []ISortable {
	for i := 1; i < len(sortables); i++ {
		el := sortables[i]
		for j := 0; j <= i-1; j++ {
			if el.Compare(sortables[j]) < 0 {
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
