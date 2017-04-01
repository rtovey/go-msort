package msort

import (
	"bufio"
	"os"
)

func distribute(unsortedFile string, es *ExternalSorter) []string {
	f, err := os.Open(unsortedFile)
	check(err)
	defer f.Close()
	unsorted := bufio.NewScanner(f)

	current := NewPriorityQueue()
	next := NewPriorityQueue()

	prefillQueue(unsorted, current, es.capacity)

	runs := createRuns(unsorted, current, next)
	return runs
}

func prefillQueue(unsorted *bufio.Scanner, queue *PriorityQueue, capacity int) {
	for queue.size() < capacity && unsorted.Scan() {
		queue.Push(unsorted.Text())
	}
}

func createRuns(unsorted *bufio.Scanner, current *PriorityQueue, next *PriorityQueue) []string {
	runs := make([]string, 1)

	for !current.IsEmpty() {
		sorted := createRun(unsorted, current, next)
		run := write(sorted)
		runs = append(runs, run)
		swap(current, next)
	}

	return runs
}

func swap(current *PriorityQueue, next *PriorityQueue) {
	tmp := current
	next = current
	current = tmp
}
