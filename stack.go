package main

import "errors"

// The stack data structure
type Stack struct {
	top *Element
	num int
}

// Elements of the stack
type Element struct {
	content float64
	next    *Element
}

// Pushes an element to the top of the stack
func (stack *Stack) push(value float64) {
	stack.top = &Element{value, stack.top}
	stack.num++
}

// Determines the stack's size aka its depth
func (stack *Stack) size() int {
	return stack.num
}

// Remove and return the topmost element from the stack
// If the stack is empty, return an error
func (stack *Stack) pop() (value float64, err error) {
	if stack.num > 0 {
		value = stack.top.content
		stack.top = stack.top.next
		stack.num--
		return value, nil
	}
	return 0, errors.New("Stack Empty")
}

// Safely extract the two topmost values of the stack
func (stack *Stack) popTwo() (v1 float64, v2 float64, err error) {
	err = nil
	v1, v2 = 0, 0
	if stack.size() > 1 {
		// we can safely ignore the errors due to the length check
		v2, _ = stack.pop()
		v1, _ = stack.pop()
	} else {
		err = errors.New("Not enough arguments")
	}
	return
}

// Construct a stack
func NewStack() *Stack {
	stack := new(Stack)
	stack.num = 0
	stack.top = nil
	return stack
}
