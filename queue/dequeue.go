package queue

import (
	"errors"
)

func (q *Queue) Dequeue() (interface{} , error) {
	
	if (len(q.items) >= 1) {
		dequeued := q.items[0]
		q.items = q.items[1:]
		return dequeued , nil
	} else {
		return nil, errors.New("Queue is empty, cannot dequeue")
	}
}