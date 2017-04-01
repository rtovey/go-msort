package msort

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write(sortedData []string) string {
	tmpfile := createTemporaryFile()

	for _, line := range sortedData {
		if _, err := tmpfile.WriteString(line + "\r\n"); err != nil {
			log.Fatal(err)
		}
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	return tmpfile.Name()
}

func createTemporaryFile() *os.File {
	tmpfile, err := ioutil.TempFile("", "run")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tmpfile.Name())
	return tmpfile
}

func cleanup(files []string) {
	for _, file := range files {
		os.Remove(file)
	}
}
