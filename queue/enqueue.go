package queue

import (
	"reflect"
	"errors"
)

var queueSize int			= 10
var queueSizeMultiplier int	= 10

func (q *Queue) Enqueue (data interface{}) (int, error) {
	if q.items == nil {
		q.items = make([]interface{} , 0, queueSize)

	} else if (reflect.TypeOf(q.items[0]) != reflect.TypeOf(data) ) {
		return len(q.items), errors.New("Passed object is of different type than queue's elements.")

	}

	q.items 	= append(q.items, data)
	return len(q.items), nil 
}


