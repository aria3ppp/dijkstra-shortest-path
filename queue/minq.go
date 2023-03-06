package queue

import "golang.org/x/exp/constraints"

type MinPriorityQueue[T any, O constraints.Ordered] struct {
	items         []T
	priorityValue func(T) O
}

func NewMinPriority[T any, O constraints.Ordered](
	priorityValue func(T) O,
) *MinPriorityQueue[T, O] {
	return &MinPriorityQueue[T, O]{
		items:         make([]T, 0),
		priorityValue: priorityValue,
	}
}

func (pq *MinPriorityQueue[T, O]) Enqueue(newItem T) {
	if len(pq.items) == 0 {
		pq.items = append(pq.items, newItem)
		return
	}
	for i, value := range pq.items {
		if pq.priorityValue(newItem) < pq.priorityValue(value) {
			pq.items = append(pq.items[:i+1], pq.items[i:]...)
			pq.items[i] = newItem
			return
		}
	}
	pq.items = append(pq.items, newItem)
}

func (pq *MinPriorityQueue[T, O]) Dequeue() (item T) {
	if len(pq.items) == 0 {
		return item
	}
	item = pq.items[0]
	pq.items = pq.items[1:]
	return item
}

func (pq *MinPriorityQueue[T, O]) Size() int {
	return len(pq.items)
}

func (pq *MinPriorityQueue[T, O]) Empty() bool {
	return pq.Size() == 0
}
