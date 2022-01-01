package main

import "fmt"

func main() {
	//s := Constructor()
	//s.Push(-2)
	//s.Push(0)
	//s.Push(-3)
	//fmt.Printf("after push stack : %v min : %d\n", s.Stack, *s.Min)
	//fmt.Printf("Get min : %d\n", s.GetMin())
	//s.Pop()
	//fmt.Printf("Top : %d\n", s.Top())
	//fmt.Printf("GetMin : %d\n", s.GetMin())
	//["MinStack", "push", "push", "push", "top","pop", "getMin", "pop", "getMin", "pop","push", "top", "getMin", "push", "top",getMin","pop","getMin"]
	//	[[], [2147483646], [2147483646], [2147483647], [], [], [], [], [], [], [2147483647], [], [], [-2147483648], [], [], [], []]
	c := Constructor()
	c.Push(2147483646)
	c.Push(2147483646)
	c.Push(2147483647)
	fmt.Printf("Top : %d\n", c.Top())
	c.Pop()
	fmt.Printf("GetMin : %d\n", c.GetMin())
	c.Pop()
	fmt.Printf("GetMin : %d\n", c.GetMin())
	c.Pop()
	c.Push(2147483647)
	fmt.Printf("Top : %d\n", c.Top())
	fmt.Printf("GetMin : %d\n", c.GetMin())
	c.Push(-2147483648)
	fmt.Printf("Top : %d\n", c.Top())
	fmt.Printf("GetMin : %d\n", c.GetMin())
	c.Pop()
	fmt.Printf("GetMin : %d\n", c.GetMin())
}

type MinStack struct {
	Stack [][]int
	Min   *int
}

func Constructor() MinStack {
	return MinStack{Stack: make([][]int, 0)}
}

// -2,-2 0,-2 -3,-3
func (this *MinStack) Push(val int) {
	var min int
	if this.Min == nil || val < *this.Min {
		this.Min = &val
		min = val
	} else {
		min = *this.Min
	}
	this.Stack = append(this.Stack, []int{val, min})
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	if len(this.Stack) != 0 {
		this.Min = &this.Stack[len(this.Stack)-1][1]
	} else {
		this.Min = nil
	}
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1][0]
}

func (this *MinStack) GetMin() int {
	return *this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
