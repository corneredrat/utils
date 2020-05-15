package main
import (
	"reflect"
	"errors"
)
func (q *Queue) Enqueue (data interface{}) (int, error) {
	if q.items == nil {
		q.items = make([]interface{} , 0, queueSize)
	
	} else if (reflect.TypeOf(q.items[0]) != reflect.TypeOf(data) ) {
		return len(q.items), errors.New("Passed object is of different type than queue.")
	}
	q.items = append(q.items, data)
	return len(q.items), nil 
}