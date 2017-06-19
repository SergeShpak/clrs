package ch2

import (
	"github.com/SergeyShpak/clrs/sorting/sorting"
)

func BinarySearch(stock []int, el int) int {
	if stock == nil {
		return -1
	}
	if len(stock) == 0 {
		return -1
	}
	sorted := binarySearchSort(stock)
	return binarySearchInternal(sorted, el)
}

func binarySearchSort(stock []int) []int {
	sorter := sorting.MergeSorter{}
	sortables := sorting.ToISortables(sorting.IntSortables(stock))
	sorter.Sort(sortables)
	return sorting.ToIntSortables(sortables)
}

func binarySearchInternal(stock []int, el int) int {
	start := 0
	end := len(stock) - 1
	for end > start {
		middle := (end + start) / 2
		if stock[middle] == el {
			return middle
		}
		if stock[middle] > el {
			end = middle - 1
			continue
		}
		start = middle + 1
	}
	if end == start {
		if stock[end] == el {
			return end
		}
	}
	return -1
}
