package testsorting

/*
import (
	"github.com/SergeyShpak/clrs/sorting/sorting"
	"testing"
)
*/

/*
func Test_SelectionSort_Sort_Sanity(t *testing.T) {
	first := sorting.IntSortable{Value: 30}
	second := sorting.IntSortable{20}
	third := sorting.IntSortable{10}
	fourth := sorting.IntSortable{0}
	sortables := []sorting.IntSortable{first, second, third, fourth}
	genericSortables := sorting.ToISortable(sortables)
	sorter := sorting.SelectionSorter{}
	sorter.Sort(genericSortables, sorting.Ascending)
	result := sorting.ISortableToIntSortable(genericSortables)
	expected := []sorting.IntSortable{fourth, third, second, first}
	for i, el := range result {
		if el.Compare(expected[i]) != 0 {
			t.Errorf("Test failed: %d != %d (index: %d)", el, expected[i], i)
		}
	}
}
*/
