package leetcode

type NiNode struct {
	val   int
	next  *NiNode
	front *NiNode
}
type MyCircularQueue struct {
	length int
	cap    int
	head   *NiNode
	tail   *NiNode
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{cap: k}

}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	n := &NiNode{val: value}
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.length++
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.head = q.head.next
	q.length--
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1

	}
	return q.head.val
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1

	}
	return q.tail.val
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.length == q.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
