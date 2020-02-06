package main

import "testing"

func Test_queue_addElementsToQueue(t *testing.T) {
	q := queue{}
	q.addElementToQueue(1)
	q.addElementToQueue(2)
	q.addElementToQueue(3)
	if q.len() != 3 {
		t.Error("Length of the queue should be 3, got: ", q.len())
	}
	firstAtQueue := q.deleteFirstFromQueue()
	if firstAtQueue!=1{
		t.Error("The first person in queue, got:", firstAtQueue)
	}
	if q.firstAtQueue()!=2{
		t.Error("If first person left queue, then second in queue will be first, got:", 2)
	}
	if q.lastAtQueue()!=3{
		t.Error("If first person left queue, then last people should stay last, and still be 3, got:", 3)
	}
	if q.len() != 2 {
		t.Error("After delete, length of the queue should be 2, got: ", q.len())
	}
}

func Test_emptyQueue(t *testing.T) {
	q := queue{}
	q.addElementToQueue(0)
	if q.len() != 1 {
		t.Error("Queue length must be 0, got: ", q.len())
	}
	if q.firstAtQueue()!=q.first.value{
		t.Error("If nobody at queue, nobody will be first, got:",q.first.value)
	}
	if q.lastAtQueue()!=q.first.value{
		t.Error("If nobody at queue, nobody will be last, got:",q.last.value)
	}
}

func Test_addToQueue(t *testing.T) {
	q := queue{}
	q.addElementToQueue(0)
	if q.len()!=1{
		t.Error("If there is one person in queue then the length of the queue is 1, got:",q.len())
	}
	if q.firstAtQueue() != q.lastAtQueue() {
		t.Error("After adding a person, length should be 1, got: ", q.len())
	}
	if q.firstAtQueue()!=q.first.value{
		t.Error("If there is one person in queue, then he’s the first, got:",q.firstAtQueue())
	}
	if q.lastAtQueue()!=q.first.value{
		t.Error("If there is one person in queue, then he is the first and last, got:",q.lastAtQueue())
	}
	q.deleteFirstFromQueue()
	if q.len() != 0 {
		t.Error("After delete length should be 0, got: ", q.len())
	}
}
