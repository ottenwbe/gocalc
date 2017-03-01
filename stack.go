//Copyright 2016 Beate Ottenwälder
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

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

// Safely extract the two topmost elements of the stack
func (stack *Stack) popTwo() (secondToTop float64, topElement float64, err error) {
	err = nil
	secondToTop, topElement = 0, 0
	if stack.size() > 1 {
		// we can safely ignore the errors due to the length check
		topElement, _ = stack.pop()
		secondToTop, _ = stack.pop()
	} else {
		err = errors.New("Not enough arguments")
	}
	return
}

// Construct a stack
func NewStack() *Stack {
	stack := &Stack{
		num: 0,
		top: nil,
	}
	return stack
}
