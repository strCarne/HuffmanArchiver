package priorityq

import "container/heap"

type PriorityObj interface {
	Priority() int
}

type PriorityQ[T PriorityObj] struct {
	heap []T
}

func NewPrQ[T PriorityObj](slice []T) *PriorityQ[T] {
	prQ := &PriorityQ[T]{}
	prQ.heap = slice
	heap.Init(prQ)
	return prQ
}

func (p *PriorityQ[T]) Len() int {
	return len(p.heap)
}

func (p *PriorityQ[T]) Less(i, j int) bool {
	return p.heap[i].Priority() > p.heap[j].Priority()
}

func (p *PriorityQ[T]) Swap(i, j int) {
	p.heap[i], p.heap[j] = p.heap[j], p.heap[i]
}

func (p *PriorityQ[T]) Push(x any) {
	p.heap = append(p.heap, x.(T))
}

func (p *PriorityQ[T]) Pop() any {
	res := p.heap[p.Len()-1]
	p.heap = p.heap[:p.Len()-1]
	return res
}
