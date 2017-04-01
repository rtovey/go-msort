package msort

import "bufio"

func createRun(unsorted *bufio.Scanner, current *PriorityQueue, next *PriorityQueue) []string {
	sortedLines := make([]string, 1)

	for !current.IsEmpty() {
		min := current.Pop()
		sortedLines = append(sortedLines, min)

		if !unsorted.Scan() {
			continue
		}

		if unsorted.Text() < min {
			next.Push(unsorted.Text())
		} else {
			current.Push(unsorted.Text())
		}
	}

	return sortedLines
}
