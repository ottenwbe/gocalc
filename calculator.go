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
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Sum of a and b
func sum(a float64, b float64) (float64, error) {
	return a + b, nil
}

// Difference between a and b
func diff(a float64, b float64) (float64, error) {
	return a - b, nil
}

// Product of a and b
func prod(a float64, b float64) (float64, error) {
	return a * b, nil
}

// Divide a by b
// Returns a division by zero error when b is 0
func div(a float64, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return 0, errors.New("Divison by zero")
	}
}

// Evaluates a postfix term and returns either the result or
// an error if the term cannot be evaluated.
//
// While the term is parsed values are pushed to a stack. This way,
// operators can be evaluated by applying the operator to the topmost elements of the stack.
func evaluate(term []string) (result float64, err error) {
	// by default, no error is returned
	err = nil
	// the default result
	result = 0.0
	// the stack which allows us to evaluate the postfix term
	stack := NewStack()
	// references to all operations
	methods := map[string]func(float64, float64) (float64, error){"+": sum, "-": diff, "/": div, "*": prod}

	// iterate over all elements of the term to evaluate it
	for _, next := range term {
		// if the next token in the term represents an operator,
		// then execute the operation on the two topmost elements of the stack
		// and push the result to the stack
		if method, ok := methods[next]; ok {
			// extract the first two values from stack
			v1, v2, tmpErr := stack.popTwo()
			if tmpErr != nil {
				err = tmpErr
				break //each error breaks the evaluation
			}
			// execute operation on the extracted values
			methodResult, tmpErr := method(v1, v2)
			if tmpErr != nil {
				err = tmpErr
				break //each error breaks the evaluation
			}
			// push result to the stack
			stack.push(methodResult)
		} else {
			// try to convert the next string in the term to a float64
			nextFloat, tmpErr := strconv.ParseFloat(next, 64)
			if tmpErr != nil {
				err = tmpErr
				break //each error breaks the evaluation
			}
			// push the next float of the term to the stack
			stack.push(nextFloat)
		}
	}

	// Extract result from stack if no error occurred and exactly one element is on the stack
	if err == nil {
		if stack.size() > 1 {
			err = errors.New("Not enough operators to evaluate complete expression.")
		} else if stack.size() == 0 {
			err = errors.New("Resultset is empty.")
		} else {
			result, err = stack.pop()
		}
	}

	return
}

func main() {

	// Get all command line arguments, except for the first one---the program.
	// The command line arguments represent the postfix term that needs to be evaluated.
	term := os.Args[1:]

	// Evaluate the postfix term
	result, err := evaluate(term)
	if err == nil {
		fmt.Println(fmt.Sprintf("%s = %f\n", term, result))
	} else {
		fmt.Println(fmt.Sprintf("Term cannot be evaluated: %s", err))
	}
}
