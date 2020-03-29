package main

import (
	"fmt"
)

type MyCircularQueue struct {
	Max int
	item [3]int
	front int
	rear int
	flag bool
}


///** Initialize your data structure here. Set the size of the queue to be k. */
//func Constructor(k int) MyCircularQueue {
//	queue := &MyCircularQueue{
//		Max: 5,
//		front:    0,
//		rear:    0,
//	}
//
//}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	//注意第一次进来，rear其实指向的是尾部加一的位置
	this.item[this.rear] = value
	this.rear = (this.rear + 1) % this.Max

	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty(){
		return false
	}
	var _ = this.item[this.front]
	this.front = (this.front+1)%this.Max

	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	return this.front
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	return this.rear
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.rear+1)%this.Max == this.front
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

//主函数
func main() {
	//初始化一个队列
	queue := &MyCircularQueue{
		Max: 3,
		front:    0,
		rear:    0,
		flag:true,
	}

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())

	fmt.Println(queue.EnQueue(1),queue.rear,queue.front)
	fmt.Println(queue.EnQueue(2),queue.rear,queue.front)
	fmt.Println(queue.EnQueue(5),queue.rear,queue.front)




}