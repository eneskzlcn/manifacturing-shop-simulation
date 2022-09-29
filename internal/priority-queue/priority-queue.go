package priority_queue

import "log"

/* This queue uses priority approach; item has little priority value is more prior. */

type QueueItem interface {
	GetPriority() int
}
type PriorityQueue[T QueueItem] struct {
	items []T
}

func NewPriorityQueue[T QueueItem]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}
func (p *PriorityQueue[T]) Enqueue(item T) {
	i := 0
	for i < len(p.items) {
		if p.items[i].GetPriority() < item.GetPriority() {
			i++
		} else {
			break
		}
	}
	if i == len(p.items) {
		p.items = append(p.items, item)
	} else {
		p.items = append(p.items[:i+1], p.items[i:]...)
		p.items[i] = item
	}
}
func (p *PriorityQueue[T]) Dequeue() QueueItem {
	item := p.items[0]
	p.items = p.items[1:]
	return item
}
func (p *PriorityQueue[T]) Print() {
	for index, item := range p.items {
		log.Printf("\n Index: %d, Item : %v", index, item)
	}
}
func (p *PriorityQueue[T]) GetItems() []T {
	return p.items
}

func (p *PriorityQueue[T]) Length() int {
	return len(p.items)
}
