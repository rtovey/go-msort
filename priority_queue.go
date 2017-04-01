package msort

import (
	"sync"
)

type PriorityQueue struct {
	items []*string
	lock  sync.Mutex
}

func NewPriorityQueue() *PriorityQueue {
	items := make([]*string, 1)
	items[0] = nil // First element of the heap queue should always be nil

	return &PriorityQueue{
		items: items,
	}
}

func (pq *PriorityQueue) Push(value string) {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	pq.items = append(pq.items, &value)
	pq.swim()
}

func (pq *PriorityQueue) Pop() string {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	if pq.size() < 1 {
		return ""
	}

	var min = *pq.items[1]

	pq.exchange(1, pq.size())
	pq.items = pq.items[0:pq.size()]
	pq.sink()

	return min
}

func (pq *PriorityQueue) IsEmpty() bool {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	return pq.size() == 0
}

func (pq *PriorityQueue) size() int {
	return len(pq.items) - 1
}

func (pq *PriorityQueue) swim() {
	k := pq.size()
	for k > 1 && pq.greater(k/2, k) {
		pq.exchange(k/2, k)
		k = k / 2
	}
}

func (pq *PriorityQueue) sink() {
	k := 1
	for 2*k <= pq.size() {
		var j = 2 * k

		if j < pq.size() && pq.greater(j, j+1) {
			j++
		}

		if !pq.greater(k, j) {
			break
		}

		pq.exchange(k, j)
		k = j
	}
}

func (pq *PriorityQueue) greater(i, j int) bool {
	return *pq.items[i] > *pq.items[j]
}

func (pq *PriorityQueue) exchange(i, j int) {
	var tempItem = pq.items[i]

	pq.items[i] = pq.items[j]
	pq.items[j] = tempItem
}
