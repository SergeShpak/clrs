package sorting

type ISorter interface {
	Sort(ISortables) ISortables
}
