package queue

//type Queue []int
// THE FIFO DATA STRUCT --- QUEUE
type Queue []interface {}

func (q *Queue) Push (v interface{}) {
	*q = append(*q, v.(int))
}

// get element from head
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}


