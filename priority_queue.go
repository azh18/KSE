package KSE

import (
	"container/heap"
	"errors"
)

const PqSize = 5


// An Item is something we manage in a priority queue.
type Item struct {
	value interface{} // The value of the item; arbitrary.
	priority float64    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewPQItem(priority float64, value interface{}) *Item {
	return &Item{
		value: value,
		priority: priority,
	}
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func NewPriorityQueue() *PriorityQueue{
	pq := make(PriorityQueue, 0)
	return &pq
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	newPriority := item.priority
	// if pq is full, and new item has larger priority, then pop the smallest to maintain top-k
	if n == PqSize && newPriority > (*pq)[0].priority{
		heap.Pop(pq)
	}
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) PopTopK(k int) ([]*Item, error){
	if k > pq.Len(){
		return nil, errors.New("there are less than k elements.\n")
	}
	ret := make([]*Item, 0)
	for i:=0;i<k;i++{
		obj := heap.Pop(pq).(*Item)
		ret = append(ret, obj)
	}
	return ret, nil
}

//// update modifies the priority and value of an Item in the queue.
//func (pq *PriorityQueue) update(item *Item, value string, priority float64) {
//	item.value = value
//	item.priority = priority
//	heap.Fix(pq, item.index)
//}