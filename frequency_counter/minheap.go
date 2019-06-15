package frequency_counter

import "container/heap"

type WordInfo struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*WordInfo

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(elem interface{}) {
	var size = len(*pq)
	var item = elem.(*WordInfo)
	item.index = size
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	var prev = *pq
	var size = len(prev)
	var item = prev[size-1]
	item.index = -1
	*pq = prev[0 : size-1]
	return item
}

func (pq *PriorityQueue) update(item *WordInfo, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type Pair struct {
	a, b interface{}
}
