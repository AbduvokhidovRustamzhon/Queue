package main

type queueNode struct {
	next     *queueNode
	prev     *queueNode
	priority int
}

type queue struct {
	first *queueNode
	last *queueNode
	size int
}

func (q *queue) len() int{
	return q.size
}

func (q *queue) addElementToQueue(queuePtr int)  {
	if q.len()==0 {
	q.first = &queueNode{
		next:     nil,
		priority: queuePtr,
	}
	q.size++
	q.last = q.first
		return
	}
	q.size++
	current := q.first
	for{
		if current.next == nil {
			current.next = &queueNode{
				next:     nil,
				priority: queuePtr,
			}
			q.last = current.next
			return
		}
		current = current.next
	}
}

func (q *queue) deleteFirstFromQueue()  queue{
	if q.len() == 0 {
		return queue{}
	}
	result := queue{
		first: q.first,
		last:  nil,
		size:  0,
	}
	q.first = q.first.next
	q.size--
	return result
}

func (q *queue) firstAtQueue()  *queueNode{
	return q.first
}

func (q *queue) lastAtQueue() *queueNode{
	return q.last
}

func main() {
}



