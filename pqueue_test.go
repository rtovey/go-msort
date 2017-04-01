package msort

import "testing"

func TestPriorityQueueItemsAreSorted(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Push("a", 3)
	pq.Push("b", 2)
	pq.Push("c", 4)

	item, priority := pq.Pop()
	assertEqual(t, "c", item)
	assertEqual(t, 4, priority)

	item, priority = pq.Pop()
	assertEqual(t, "a", item)
	assertEqual(t, 3, priority)

	item, priority = pq.Pop()
	assertEqual(t, "b", item)
	assertEqual(t, 2, priority)
}

func TestPriorityQueueIsEmptyOnInitilisation(t *testing.T) {
	pq := NewPriorityQueue()

	if !pq.IsEmpty() {
		t.Error("PriorityQueue is not empty when first initialised")
	}
}

func TestPriorityQueueIsNotEmptyAfterPush(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Push("a", 1)

	if pq.IsEmpty() {
		t.Error("PriorityQueue is empty with one item")
	}
}

func TestPriorityQueueIsEmptyAfterAllItemsPopped(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Push("a", 1)
	pq.Pop()

	if !pq.IsEmpty() {
		t.Error("PriorityQueue is not empty after all items popped")
	}
}

func TestPriorityQueueReturnsNilItemWhenEmpty(t *testing.T) {
	pq := NewPriorityQueue()

	item, _ := pq.Pop()

	if item != nil {
		t.Error("Empty PriorityQueue does not return a nil item when popped")
	}
}

func TestPriorityQueueReturnsZeroPriorityWhenEmpty(t *testing.T) {
	pq := NewPriorityQueue()

	_, priority := pq.Pop()

	if priority != 0 {
		t.Error("Empty PriorityQueue does not return 0 priority when popped")
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("%v != %v", a, b)
	}
}
