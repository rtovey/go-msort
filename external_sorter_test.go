package msort

import "testing"

func TestExternalSorterSort(t *testing.T) {
	unsortedFile := "C:\\git\\unsorted.txt"

	es := NewExternalSorter(3, 3)
	es.Sort(unsortedFile)
}
