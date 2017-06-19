package testsorting

import (
	"errors"
	"fmt"
	"github.com/SergeyShpak/clrs/sorting/sorting"
	"sort"
	"testing"
)

func InternalTest_SortSanity(sorter sorting.ISorter) error {
	sortables := sorting.CreateIntSortables(31, 41, 59, 26, 41, 58)
	genericSortables := sorting.ToISortables(sortables)
	sorter.Sort(genericSortables)
	result := sorting.ToIntSortables(genericSortables)
	expected := sortables.Copy()
	copy(sortables, expected)
	sort.Sort(expected)
	for i, el := range result {
		if el != expected[i] {
			errMsg := fmt.Sprintf("Test failed: Actual %d != Expected %d (index: %d)",
				el, expected[i], i)
			err := errors.New(errMsg)
			return err
		}
	}
	return nil
}

func Test_InsertionSort_Sort_Sanity(t *testing.T) {
	sorter := sorting.InsertionSorter{}
	err := InternalTest_SortSanity(sorting.ISorter(sorter))
	if err != nil {
		t.Error(err)
	}
}

func Test_SelectionSort_Sort_Sanity(t *testing.T) {
	sorter := sorting.SelectionSorter{}
	err := InternalTest_SortSanity(sorting.ISorter(sorter))
	if err != nil {
		t.Error(err)
	}
}

func Test_MergeSort_Sort_Sanity(t *testing.T) {
	sorter := sorting.MergeSorter{}
	err := InternalTest_SortSanity(sorting.ISorter(sorter))
	if err != nil {
		t.Error(err)
	}
}
