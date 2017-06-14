package sorting

type Order int

const (
	Ascending Order = iota
	Descending
)

type ISorter interface {
	Sort([]*ISortable, Order) []*ISortable
}
