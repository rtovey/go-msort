package msort

type item struct {
	value    interface{}
	priority int
}

type PriorityQueue struct {
	items []*item
}

func newItem(value interface{}, priority int) *item {
	return &item{
		value:    value,
		priority: priority,
	}
}

func NewPriorityQueue() *PriorityQueue {
	items := make([]*item, 1)
	items[0] = nil // First element of the heap queue should always be nil

	return &PriorityQueue{
		items: items,
	}
}

func (pq *PriorityQueue) Push(value interface{}, priority int) {
	item := newItem(value, priority)

	pq.items = append(pq.items, item)
	pq.swim()
}

func (pq *PriorityQueue) Pop() (interface{}, int) {
	if pq.size() < 1 {
		return nil, 0
	}

	var max *item = pq.items[1]

	pq.exchange(1, pq.size())
	pq.items = pq.items[0:pq.size()]
	pq.sink()

	return max.value, max.priority
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.size() == 0
}

func (pq *PriorityQueue) size() int {
	return len(pq.items) - 1
}

func (pq *PriorityQueue) swim() {
	k := pq.size()
	for k > 1 && pq.less(k/2, k) {
		pq.exchange(k/2, k)
		k = k / 2
	}
}

func (pq *PriorityQueue) sink() {
	k := 1
	for 2*k <= pq.size() {
		var j int = 2 * k

		if j < pq.size() && pq.less(j, j+1) {
			j++
		}

		if !pq.less(k, j) {
			break
		}

		pq.exchange(k, j)
		k = j
	}
}

func (pq *PriorityQueue) less(i, j int) bool {
	return pq.items[i].priority < pq.items[j].priority
}

func (pq *PriorityQueue) exchange(i, j int) {
	var tempItem *item = pq.items[i]

	pq.items[i] = pq.items[j]
	pq.items[j] = tempItem
}
