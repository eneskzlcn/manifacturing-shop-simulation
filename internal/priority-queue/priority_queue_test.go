package priority_queue_test

import (
	priority_queue "github.com/eneskzlcn/manifacturing-shop-simulation/internal/priority-queue"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type TestQueueItem struct {
	value int
}

func (t TestQueueItem) GetPriority() int {
	return t.value
}
func TestPriorityQueue(t *testing.T) {
	testItems := []TestQueueItem{
		{value: 2}, {value: 3}, {value: 20},
	}
	queue := priority_queue.NewPriorityQueue()
	assert.NotNil(t, queue)
	log.Printf("%v", testItems[0])
	queue.Enqueue(testItems[0])
	queue.Enqueue(testItems[1])
	queue.Enqueue(testItems[2])

	queueItems := queue.GetItems()
	assert.Equal(t, len(queueItems), len(testItems))

	queue.Print()
	assert.Equal(t, testItems[0].value, queueItems[0].GetPriority())
	assert.Equal(t, testItems[1].value, queueItems[1].GetPriority())

	item1 := queue.Dequeue()
	queue.Print()
	item2 := queue.Dequeue()
	item3 := queue.Dequeue()

	assert.Equal(t, item1, testItems[0])
	assert.Equal(t, item2, testItems[1])
	assert.Equal(t, item3, testItems[2])
}
