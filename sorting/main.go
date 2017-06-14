package main

import (
	"fmt"
	"github.com/SergeyShpak/clrs/sorting/sorting"
)

func main() {
	first := sorting.IntSortable{30}
	second := sorting.IntSortable{20}
	third := sorting.IntSortable{10}
	sortables := []sorting.IntSortable{first, second, third}
	genericSortables := sorting.ToISortable(sortables)
	sorter := sorting.InsertionSorter{}
	sorter.Sort(genericSortables, sorting.Ascending)
	fmt.Println(genericSortables)
}
