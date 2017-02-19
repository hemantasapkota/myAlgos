package queue

type Queue struct {
	queue []interface{}
	len   int
}

func New() *Queue {
	queue := &Queue{queue: make([]interface{}, 0), len: 0}
	return queue
}

func (q *Queue) Enqueue(val interface{}) {
	q.queue = append(q.queue, val)
	q.len++
}

func (q *Queue) Dequeue() interface{} {
	// Remove the head
	tmp := q.queue[0]
	q.queue = q.queue[1:]
	q.len--
	return tmp
}

func (q *Queue) Peek() interface{} {
	if q != nil {
		return q.queue[0]
	}
	return nil
}

func (q *Queue) Len() int {
	return q.len
}
