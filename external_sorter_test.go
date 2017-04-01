package msort

import (
	"os"
	"testing"
)

func TestExternalSorterSort(t *testing.T) {
	unsortedFile := createTemporaryFile().Name()
	defer os.Remove(unsortedFile)

	es := NewExternalSorter(3, 3)
	es.Sort(unsortedFile)
}

func createUnsortedTestFile() *os.File {
	tmpfile := createTemporaryFile()
	tmpfile.WriteString("g\r\na\r\nl\r\nd\r\nu\r\nz\r\nc\r\nx")
	tmpfile.Close()
	return tmpfile
}
