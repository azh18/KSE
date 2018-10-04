package dal

import (
	"container/heap"
	"errors"
)


// An PQItem is something we manage in a priority queue.
type PQItem struct {
	value interface{} // The value of the item; arbitrary.
	priority float64    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewPQItem(priority float64, value interface{}) *PQItem {
	return &PQItem{
		value: value,
		priority: priority,
	}
}

func (i *PQItem) GetValue() interface{}{
	return i.value
}

func (i *PQItem) GetPriority() float64{
	return i.priority
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	Content []*PQItem
	PqSize int
	IsMinHeap bool
}

func NewPriorityQueue(pqSize int, isMinHeap bool) *PriorityQueue{
	pq := &PriorityQueue{
		Content:make([]*PQItem, 0),
		PqSize: pqSize,
		IsMinHeap: isMinHeap,
	}
	return pq
}

func (pq PriorityQueue) Len() int { return len(pq.Content) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq.IsMinHeap{
		return pq.Content[i].priority < pq.Content[j].priority
	} else {
		return pq.Content[i].priority > pq.Content[j].priority
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.Content[i], pq.Content[j] = pq.Content[j], pq.Content[i]
	pq.Content[i].index = i
	pq.Content[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.Content)
	item := x.(*PQItem)
	item.index = n
	newPriority := item.priority
	// if pq is full, and new item has larger priority, then pop the smallest to maintain top-k
	if n == pq.PqSize && newPriority <= pq.Content[0].priority{
		return
	}
	if n == pq.PqSize {
		heap.Pop(pq)
	}
	pq.Content = append(pq.Content, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old.Content)
	item := old.Content[n-1]
	item.index = -1 // for safety
	pq.Content = old.Content[0 : n-1]
	return item
}

func (pq *PriorityQueue) PopTopK(k int) ([]*PQItem, error){
	if k > pq.Len(){
		return nil, errors.New("there are less than k elements.\n")
	}
	ret := make([]*PQItem, 0)
	for i:=0;i<k;i++{
		obj := heap.Pop(pq).(*PQItem)
		ret = append(ret, obj)
	}
	return ret, nil
}

//// update modifies the priority and value of an PQItem in the queue.
//func (pq *PriorityQueue) update(item *PQItem, value string, priority float64) {
//	item.value = value
//	item.priority = priority
//	heap.Fix(pq, item.index)
//}