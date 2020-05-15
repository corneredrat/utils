package main
import (
	"fmt"
)

var queueSize int32				= 10
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
	len, _ 	:= q.Enqueue(1)
	fmt.Printf("\n%v %v,%v %p \n",q.items, len, cap(q.items), &q.items)
	len, _ 	= q.Enqueue(2)
	fmt.Printf("\n%v %v,%v %p \n",q.items, len, cap(q.items), &q.items)
	len, _ 	= q.Enqueue(1)
	test	:= make([]int,10,10)
	fmt.Printf("\n%v %v,%v %p \n",q.items, len, cap(q.items), &q.items)
	fmt.Printf("%v %v, %p \n",q.items, test, &test)
	len, _ = q.Enqueue(2)
	test1	:= make([]int,10,10)
	fmt.Printf("\n%v %v,%v %p \n",q.items, len, cap(q.items), &q.items)
	fmt.Printf("%v %v, %p \n",test1, len, &test1)
}





