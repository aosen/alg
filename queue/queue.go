package queue

import (
	"errors"
	"sync"
)

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
	//队列大小
	size int
}

func NewQueue(size int) *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)
	queue.size = size
	return queue
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len
}

func (q *Queue) Empty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}

func (q *Queue) EnQueue(el interface{}) (err error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.len >= q.size {
		err = errors.New("the queue will overflow")
	} else {
		q.queue = append(q.queue, el)
		q.len++
	}
	return
}

func (q *Queue) OutQueue() (el interface{}, err error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.len == 0 {
		err = errors.New("the queue is empty")
	} else {
		el, q.queue = q.queue[0], q.queue[1:]
		q.len--
	}
	return
}

func (q *Queue) Peek() (el interface{}, err error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.len == 0 {
		err = errors.New("the queue is empty")
	} else {
		el = q.queue[0]
	}
	return
}

//阻塞队列，当数据超出队列空间大小会发生阻塞
type BlockQueue struct {
	queue chan *interface{}
}

func NewBlockQueue(size int) *BlockQueue {
	ch := make(chan *interface{}, size)
	return &BlockQueue{
		queue: ch,
	}
}

func (self *BlockQueue) EnQueue(el *interface{}) {
	self.queue <- el
}

func (self *BlockQueue) OutQueue() (el *interface{}) {
	return <-self.queue
}

func (self *BlockQueue) Len() int {
	return len(self.queue)
}
