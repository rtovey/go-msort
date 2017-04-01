package msort

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestExternalSorterSort(t *testing.T) {
	unsortedFile := createTemporaryFile().Name()
	defer os.Remove(unsortedFile)

	es := NewExternalSorter(3, 3)
	sortedFile := es.Sort(unsortedFile)

	printSortedFileContents(sortedFile)
}

func createUnsortedTestFile() *os.File {
	tmpfile := createTemporaryFile()
	tmpfile.WriteString("g\na\nl\nd\nu\nz\nc\nx")
	tmpfile.Close()
	return tmpfile
}

func printSortedFileContents(sortedFileName string) {
	data, err := ioutil.ReadFile(sortedFileName)
	check(err)
	fmt.Println(data)
}
