package main

import (
	"fmt"
	"github.com/SergeyShpak/clrs/sorting/sorting"
)

func main() {
	first := sorting.IntSortable{Value: 30}
	second := sorting.IntSortable{20}
	third := sorting.IntSortable{10}
	sortables := []sorting.IntSortable{first, second, third}
	genericSortables := sorting.ConvertToISortable(sortables)
	sorter := sorting.InsertionSorter{}
	sorter.Sort(genericSortables)
	fmt.Println(genericSortables)
}
