package msort

import "testing"

func TestPriorityQueueItemsAreSorted(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Push("z")
	pq.Push("x")
	pq.Push("y")

	assertEqual(t, "x", pq.Pop())
	assertEqual(t, "y", pq.Pop())
	assertEqual(t, "z", pq.Pop())
}

func TestPriorityQueueIsEmptyOnInitilisation(t *testing.T) {
	pq := NewPriorityQueue()

	if !pq.IsEmpty() {
		t.Error("PriorityQueue is not empty when first initialised")
	}
}

func TestPriorityQueueIsNotEmptyAfterPush(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Push("a")

	if pq.IsEmpty() {
		t.Error("PriorityQueue is empty with one item")
	}
}

func TestPriorityQueueIsEmptyAfterAllItemsPopped(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Push("a")
	pq.Pop()

	if !pq.IsEmpty() {
		t.Error("PriorityQueue is not empty after all items popped")
	}
}

func TestPriorityQueueReturnsEmptyStringWhenEmpty(t *testing.T) {
	pq := NewPriorityQueue()

	item := pq.Pop()

	if item != "" {
		t.Error("Empty PriorityQueue does not return a nil item when popped")
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("%v != %v", a, b)
	}
}
