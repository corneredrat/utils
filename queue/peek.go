package queue 

import (
	"errors"
)

func (q *Queue) PeekOldest() (interface{} , error) {
	if (len(q.items) >= 1) {
		element := q.items[0]
		return element, nil
	} else {
		return nil, errors.New()
	}
}

func (q *Queue) PeekNewest() (interface{} , error) {
	if (len(q.items) >= 1) {
		len 	:= len(q.items)
		element := q.items[len-1]
		return element, nil
	} else {
		return nil, errors.New()
	}
}
