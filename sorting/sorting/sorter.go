package sorting

type ISorter interface {
	Sort(sortables []*ISortable) []*ISortable
}
