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

import (
	"testing"
)


func TestEmptyStackPop(t *testing.T) {
	stack := new (Stack)
	_,err := stack.pop()
	if (err == nil) {
		t.Error("Expected error, however, no error occurred")
	}
}

func TestStackPushPop(t *testing.T) {
	stack := new (Stack)
	testVal := 5.0
	stack.push(testVal)
	resultVal,_ := stack.pop()
	if (resultVal != testVal) {
		t.Error("Expected error, however, no error occurred")
	}
}

func TestStackSize(t *testing.T) {
	stack := new (Stack)
	testVal := 5.0
	stack.push(testVal)
	if (stack.size() != 1) {
		t.Error("Length of stack was not correctly calculated")
	}
}

func TestEmptyStack(t *testing.T) {
	_,_,err := new(Stack).popTwo()
	if err == nil {
		t.Error("Expected error, however, no error occurred")
	}
}