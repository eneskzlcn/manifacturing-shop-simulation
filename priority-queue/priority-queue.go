package priority_queue

import "log"

/* This queue uses priority approach; item has little priority value is more prior.
*/

type QueueItem interface {
	GetPriority() int
}
type PriorityQueue struct {
	items []QueueItem
}
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}
func (p *PriorityQueue) Enqueue(item QueueItem) {
	i := 0
	for i < len(p.items) {
		if p.items[i].GetPriority() < item.GetPriority() {
			i++
		}else {
			break
		}
	}
	if i == len(p.items) {
		p.items = append(p.items, item)
	}else {
		p.items = append(p.items[:i+1],p.items[i:]...)
		p.items[i] = item
	}
}
func (p *PriorityQueue) Dequeue() QueueItem {
	item := p.items[0]
	p.items = p.items[1:]
	return item
}
func (p *PriorityQueue) Print() {
	for index, item := range p.items {
		log.Printf("\n Index: %d, Item : %v", index, item)
	}
}
func (p *PriorityQueue) GetItems() []QueueItem {
	return p.items
}

func (p *PriorityQueue) Length() int {
	return len(p.items)
}