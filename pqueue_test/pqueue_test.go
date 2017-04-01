package pqueue_test

import (
	"fmt"
	"go-msort/pqueue"
	"testing"
)

func TestItemsAddedToPriorityQueueAreSorted(t *testing.T) {
	items := map[string]int{
		"a": 3, "b": 2, "c": 4,
	}

	pq := pqueue.NewPriorityQueue()

	for key, value := range items {
		pq.Push(key, value)
	}

	fmt.Println(pq.Pop())
}
