package msort

import (
	"fmt"
	"testing"
)

func TestItemsAddedToPriorityQueueAreSorted(t *testing.T) {
	items := map[string]int{
		"a": 3, "b": 2, "c": 4,
	}

	pq := NewPriorityQueue()

	for key, value := range items {
		pq.Push(key, value)
	}

	fmt.Println(pq.Pop())
}
