package main
import (
	"fmt"
	"reflect"
	"errors"
)

var queueSize int32				= 1
var queueSizeMultiplier int16	= 10
var head int32					= 0
var tail int32					= 0
type element struct {
	value interface{}
}

type Queue struct {
	items []interface{}
}

func main () {
	q := Queue {}
	len, _ := q.Enqueue(1)
	fmt.Printf("%v %v, %p \n",q.items, len, &q.items)
	len, _ = q.Enqueue(2)
	fmt.Printf("%v %v, %p \n",q.items, len, &q.items)
	len, _ = q.Enqueue(1)
	fmt.Printf("%v %v, %p \n",q.items, len, &q.items)
	len, _ = q.Enqueue(2)
	fmt.Printf("%v %v, %p \n",q.items, len, &q.items)
}

func (q *Queue) Enqueue (data interface{}) (int, error) {
	if q.items == nil {

		q.items = make([]interface{} , 0, queueSize)
				
	
	} else if (reflect.TypeOf(q.items[0]) != reflect.TypeOf(data) ) {
		return len(q.items), errors.New("Passed object is of different type than queue.")
	}
	q.items = append(q.items, data)
	return len(q.items), nil 
}

