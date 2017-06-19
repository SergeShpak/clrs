package sorting

type MergeSorter struct{}

func (sorter MergeSorter) Sort(sortables ISortables) ISortables {
	return sorter.mergeSort(sortables, 0, sortables.Len()-1)
}

func (sorter MergeSorter) mergeSort(sortables ISortables, start int,
	end int) ISortables {
	if start >= end {
		return sortables
	}
	middle := (end + start) / 2
	sorter.mergeSort(sortables, start, middle)
	sorter.mergeSort(sortables, middle+1, end)
	sorter.merge(sortables, start, middle, end)
	return sortables
}

func (sorter MergeSorter) merge(sortables ISortables, start int, middle int,
	end int) ISortables {
	stock := make([]interface{}, (end - start + 1))
	currLeft := start
	currRight := middle + 1
	stockInd := 0
	for currLeft <= middle || currRight <= end {
		if currLeft <= middle && (currRight > end ||
			sortables.Less(currLeft, currRight)) {
			stock[stockInd] = sortables.Get(currLeft)
			currLeft++
			stockInd++
			continue
		}
		stock[stockInd] = sortables.Get(currRight)
		currRight++
		stockInd++
	}
	for i := 0; i < len(stock); i++ {
		sortables.Set(start+i, stock[i])
	}
	return sortables
}
