package msort

type ExternalSorter struct {
	capacity   int
	mergeCount int
}

func NewExternalSorter(capacity int, mergeCount int) *ExternalSorter {
	return &ExternalSorter{
		capacity:   capacity,
		mergeCount: mergeCount,
	}
}

func (es *ExternalSorter) Sort(unsortedFile string) string {
	runs := distribute(unsortedFile, es)
	sorted := merge(runs)
	cleanup(runs)
	return sorted
}
