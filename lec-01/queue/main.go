package main

import "fmt"

// queue
type Queue struct {
	array []int
}

// EnQueue
func (q *Queue) EnQueue(i int) {
	q.array = append(q.array, i)
}

// DeQueue
func (q *Queue) Dequeue() int {
	toRemove := q.array[0]
	q.array = q.array[1:]
	return toRemove
}

func main() {
	queue := Queue{}
	fmt.Println(queue)
	queue.EnQueue(100)
	queue.EnQueue(200)
	queue.EnQueue(300)
	fmt.Println(queue)
	queue.Dequeue()
	fmt.Println(queue)
}
