package main

import "fmt"

type MinStack struct {
	item []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	a := MinStack{
		make([]int,0),
		make([]int,1),
	}
	a.min[0] = 0
	return a
}

func (this *MinStack) Push(x int)  {
	this.item = append(this.item,x)

	if x <= this.min[0]{
		this.min = append(this.min, x)
		for i := 0; i < len(this.min); i++{
			for j := i+1; j < len(this.min); j++{
				if j > i{
					this.min[j],this.min[i] = this.min[i],this.min[j]
				}
			}
		}
	}
}

func (this *MinStack) Pop()  {
	item := this.Top()

	this.item = this.item[:len(this.item) - 1]

	if item <= this.GetMin(){
		this.min[0] = 0
	}
}


func (this *MinStack) Top() int {
	if len(this.item) == 0{
		panic("is nil")
	}
	item := this.item[len(this.item) - 1]
	return item
}


func (this *MinStack) GetMin() int {
	if len(this.min) == 0{
		panic("is nil")
	}
	min := this.min[len(this.min) - 1]
	return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main()  {
	a := Constructor()
	a.Push(-2)
	a.Push(0)
	a.Push(-3)
	fmt.Println(a.GetMin())
	a.Pop()
	fmt.Println(a.Top())
	fmt.Println(a.GetMin())
	//minStack.push(0);
	//minStack.push(-3);
	//minStack.getMin();   --> 返回 -3.
	//minStack.pop();
	//minStack.top();      --> 返回 0.
	//minStack.getMin();
}