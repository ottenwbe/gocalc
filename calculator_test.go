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

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    float64
		wantErr bool
	}{
		{"sum", []string{"1", "1", "+"}, 2, false},
		{"diff", []string{"3", "2", "-"}, 1, false},
		{"div", []string{"22", "2", "/"}, 11, false},
		{"prod", []string{"2", "5", "*"}, 10, false},
		{"div by zero", []string{"22", "0", "/"}, 0, true},
		{"one value", []string{"22", "/"}, 0, true},
		{"invalid input", []string{"22", "a", "/"}, 0, true},
		{"empty input", []string{}, 0, true},
		{"only numbers", []string{"1", "1"}, 0, true},
		// multi-operator term: (2 + 3) * 4 = 20
		{"multi-op", []string{"2", "3", "+", "4", "*"}, 20, false},
		// negative intermediate result: 3 - 5 = -2
		{"negative", []string{"3", "5", "-"}, -2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := evaluate(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.want {
				t.Errorf("got %v, want %v", result, tt.want)
			}
		})
	}
}
