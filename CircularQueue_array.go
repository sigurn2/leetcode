//package leetcode
//
//type MyCircularQueue struct {
//	front  int
//	rear   int
//	target []int
//}
//
//func Constructor(k int) MyCircularQueue {
//	return MyCircularQueue{target: make([]int, k+1)}
//}
//
//func (q *MyCircularQueue) EnQueue(value int) bool {
//	if q.IsFull() {
//		return false
//	} else {
//		q.target[q.rear] = value
//		q.rear = (q.rear + 1) % len(q.target)
//		return true
//	}
//}
//
//func (q *MyCircularQueue) DeQueue() bool {
//	if q.IsEmpty() {
//		return false
//	} else {
//		q.front = (q.front + 1) % len(q.target)
//		return true
//	}
//}
//
//func (this *MyCircularQueue) Front() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	return this.target[this.front]
//}
//
//func (this *MyCircularQueue) Rear() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	return this.target[(this.rear-1+len(this.target))%len(this.target)]
//}
//
//func (this *MyCircularQueue) IsEmpty() bool {
//	return this.rear == this.front
//}
//
//func (this *MyCircularQueue) IsFull() bool {
//	return (this.rear+1)%len(this.target) == this.front
//}
//
