package minStack

import "math"

type MinStack struct {
	stack []int
	// 存当前入栈元素的最小值，弹出的时候一并弹出
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	top := s.minStack[len(s.minStack)-1]
	s.minStack = append(s.minStack, min(x, top))
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
