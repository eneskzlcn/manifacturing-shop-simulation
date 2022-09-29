package priority_queue_test

import (
	"github.com/eneskzlcn/manufacturing-shop-simulation/internal/priority-queue"
	"github.com/stretchr/testify/assert"
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
	queue := priority_queue.NewPriorityQueue[TestQueueItem]()
	assert.NotNil(t, queue)

	queue.Enqueue(testItems[0])
	queue.Enqueue(testItems[1])
	queue.Enqueue(testItems[2])

	queueItems := queue.GetItems()
	assert.Equal(t, len(queueItems), len(testItems))

	assert.Equal(t, testItems[0].value, queueItems[0].GetPriority())
	assert.Equal(t, testItems[1].value, queueItems[1].GetPriority())

	item1 := queue.Dequeue()

	item2 := queue.Dequeue()
	item3 := queue.Dequeue()

	assert.Equal(t, item1, testItems[0])
	assert.Equal(t, item2, testItems[1])
	assert.Equal(t, item3, testItems[2])
}
