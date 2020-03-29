package main

import "fmt"

type Element interface{}

type Queue interface {

	Offer(e Element)    //向队列中添加元素
	Poll()   Element    //移除队列中最前面的元素
	Clear()  bool       //清空队列
	Size()  int            //获取队列的元素个数
	IsEmpty() bool   //判断队列是否是空
}

type  sliceEntry struct{
	element []Element
}


func NewQueue() *sliceEntry {
	return &sliceEntry{}
}

func (entry *sliceEntry) Offer(e Element)  {
	entry.element = append(entry.element,e)
}


func (entry *sliceEntry) Poll() Element{
	if entry.IsEmpty() {
		fmt.Println("queue is nil")
		return nil
	}

	fnode := entry.element[0]
	entry.element = entry.element[1:]

	return fnode
}

func (entry *sliceEntry) Clear() bool {
	if entry.IsEmpty() {
		fmt.Println("queue is nil")
		return false
	}

	for i:= 0;i < len(entry.element);i++{
		entry.element[i] =nil
	}

	entry.element = nil
	return true
}

func (entry *sliceEntry)Size() int {
	return len(entry.element)
}


func (entry *sliceEntry)IsEmpty() bool {
	if len(entry.element) == 0{
		return true
	}
	return false
}


func main() {
	queue := NewQueue()
	for i := 0; i < 50; i++ {
		queue.Offer(i)
	}


	fmt.Println(queue)
	fmt.Println("size:", queue.Size())
	fmt.Println("移除最前面的元素：", queue.Poll())
	fmt.Println("size:", queue.Size())
	for i := 0; i < 50; i++ {
		queue.Offer(i)
	}
	fmt.Println(queue)
	fmt.Println("清空：", queue.Clear())
	for i := 0; i < 50; i++ {
		queue.Offer(i)
	}
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
	fmt.Println(queue.Size())
}