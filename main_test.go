package main

import "testing"

func Test_queue_addElementsToQueue(t *testing.T) {
	q := queue{}
	q.addElementToQueue(1)
	q.addElementToQueue(2)
	q.addElementToQueue(3)
	if q.len() != 3{
		t.Error("length of the queue should be 3, got: ", q.len())
	}
	q.deleteFirstFromQueue()
	if q.len() != 2{
		t.Error("after delete, length of the queue should be 2, got: ", q.len())
	}
}

func Test_emptyQueue(t *testing.T) {
	q := queue{}
	q.addElementToQueue(0)
	if q.len()!= 1 {
		t.Error("length should be 0, got: ", q.len())
	}
}

func Test_addToQueue(t *testing.T) {
	q := queue{}
	q.addElementToQueue(0)
	if q.firstAtQueue() != q.lastAtQueue() {
		t.Error("after adding a person, length should be 1, got: ", q.len())
	}
	q.deleteFirstFromQueue()
	if q.len() != 0 {
		t.Error("after delete length should be 0, got: ", q.len())
	}
}