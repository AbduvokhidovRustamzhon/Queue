package main

type queueNode struct {
	next  *queueNode
	prev  *queueNode
	value interface{}
}

type queue struct {
	first *queueNode
	last  *queueNode
	size  int
}

func (q *queue) len() int {
	return q.size
}

func (q *queue) addElementToQueue(value int) {
	if q.len() == 0 {
		q.first = &queueNode{
			next:  nil,
			value: value,
		}
		q.size++
		q.last = q.first
		return
	}
	q.size++
	current := q.first
	for {
		if current.next == nil {
			current.next = &queueNode{
				next:  nil,
				value: value,
			}
			q.last = current.next
			return
		}
		current = current.next
	}

}

func (q *queue) deleteFirstFromQueue() interface{} {
	if q.len() == 0 {
		return nil
	}
	result := q.first.value

	q.first = q.first.next
	q.size--

	return result
}

func (q *queue) firstAtQueue() interface{} {
	if q.first == nil {
		return nil
	}
	return q.first.value
}

func (q *queue) lastAtQueue() interface{} {
	if q.last == nil {
		return nil
	}
	return q.last.value
}

func main() {
}
