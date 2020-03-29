package main

type MyStack struct {
    sli []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
    this.sli = append(this.sli, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    temp := this.sli[len(this.sli) - 1]
    this.sli = this.sli[:len(this.sli) - 1]
    return temp
}

/** Get the top element. */
func (this *MyStack) Top() int {
    return this.sli[len(this.sli) - 1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.sli) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
