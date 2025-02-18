package datastruc

import "errors"

type Element struct {
	Value    int
	Previous *Element
}

type Queue struct {
	Length   int
	Capacity int
	First    *Element
	Last     *Element
}

func (q *Queue) Enqueue(data int) (int, error) {
	if q.IsFull() {
		return -1, errors.New("queue is full")
	}

	elt := Element{
		Value:    data,
		Previous: nil,
	}
	q.Last = &elt
	q.Length++

	if q.IsEmpty() {
		q.First = &elt
	}

	return q.Last.Value, nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("empty queue")
	}
	output := q.First.Value
	q.First = q.First.Previous
	q.Length--
	return output, nil
}

func (q *Queue) Size() int {
	return q.Length
}

func (q *Queue) Peek() int {
	return q.First.Value
}

func (q *Queue) IsEmpty() bool {
	return q.Size() <= 0
}

func (q *Queue) IsFull() bool {
	return q.Size() >= q.Capacity
}
