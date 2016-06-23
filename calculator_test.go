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

func TestBasicSum(t *testing.T) {
	input := []string{"1", "1", "+"}
	result, _ := evaluate(input)
	if result != 2 {
		t.Error("Expected 2 got ", result)
	}
}

func TestBasicDiff(t *testing.T) {
	input := []string{"3", "2", "-"}
	result, _ := evaluate(input)
	if result != 1 {
		t.Error("Expected 1 got ", result)
	}
}

func TestBasicDiv(t *testing.T) {
	input := []string{"22", "2", "/"}
	result, _ := evaluate(input)
	if result != 11 {
		t.Error("Expected 11 got ", result)
	}
}

func TestBasicProd(t *testing.T) {
	input := []string{"2", "5", "*"}
	result, _ := evaluate(input)
	if result != 10 {
		t.Error("Expected 10 got ", result)
	}
}

func TestDivByZero(t *testing.T) {
	input := []string{"22", "0", "/"}
	_,err := evaluate(input)
	if err == nil {
		t.Error("Expected error, however, no error occurred")
	}
}

func TestOneVal(t *testing.T) {
	input := []string{"22", "/"}
	_,err := evaluate(input)
	if err == nil {
		t.Error("Expected error, however, no error occurred")
	}
}

func TestInvalidInput(t *testing.T) {
	input := []string{"22", "a", "/"}
	_,err := evaluate(input)
	if err == nil {
		t.Error("Expected error, however, no error occurred")
	}
}